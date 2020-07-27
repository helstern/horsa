package process

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func outputWriterFactory(src io.ReadCloser, dest io.Writer) func() error {
	return func() error {
		_, err := io.Copy(dest, src)
		closeError := src.Close() // in case io.Copy stopped due to write error
		if closeError == nil {
			log.Print("closed an output reader")
		}
		return err
	}
}

func inputWriterFactory(src io.Reader, dest io.WriteCloser) func() error {
	return func() error {
		_, err := io.Copy(dest, src)
		closeError := dest.Close() // in case io.Copy stopped due to write error
		if closeError == nil {
			log.Print("closed an input writer")
		}
		return err
	}
}

func waitForProcess(done chan<- error, proc *os.Process, procStartError error) {

	if procStartError != nil {
		done <- procStartError
		return
	}

	var err error

	_, err = proc.Wait()
	done <- err
}

func HandleConnection(ctx context.Context, conn net.Conn, args []string) error {

	var err error

	// the proxy pipes of our process
	stdinR, stdinW, _ := os.Pipe()
	stdoutR, stdoutW, _ := os.Pipe()
	stderrR, stderrW, _ := os.Pipe()

	// start the goroutines that write from and to the process' pipes
	ioWriters := [](func() error){
		outputWriterFactory(stdoutR, conn),
		outputWriterFactory(stderrR, conn),
		inputWriterFactory(conn, stdinW),
	}

	onDone := make(chan error)

	for _, ioWritter := range ioWriters {
		go ioWritter()
	}

	var proc *os.Process
	proc, err = os.StartProcess(args[0], args, &os.ProcAttr{
		Files: []*os.File{stdinR, stdoutW, stderrW},
		Env:   os.Environ(),
		//Sys:   &syscall.SysProcAttr{
		//	Noctty: false,
		//},
	})
	go waitForProcess(onDone, proc, err)

	select {
	case <-onDone:
	case <-ctx.Done():
		proc.Signal(os.Interrupt)
		proc.Kill()

	}
	// close pipes
	_ = stdoutW.Close()
	_ = stderrW.Close()
	_ = stdinR.Close()

	if err != nil {
		return fmt.Errorf("error starting process: %w", err)
	}

	// should try to see if release does something
	//proc.Release()

	//var procState *os.ProcessState
	//procState, err = proc.Wait()

	err = conn.Close()
	if err != nil {
		return fmt.Errorf("error closing the connection: %w", err)
	}
	return err
}
