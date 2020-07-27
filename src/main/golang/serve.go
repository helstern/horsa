package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
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

func serve(conn net.Conn, args []string) {

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
	for _, ioWritter := range ioWriters {
		go ioWritter()
	}

	log.Print(args[0])
	log.Print(strings.Join(args[1:], ","))

	var proc *os.Process
	proc, err = os.StartProcess(args[0], args, &os.ProcAttr{
		Files: []*os.File{stdinR, stdoutW, stderrW},
		Env:   os.Environ(),
		//Sys:   &syscall.SysProcAttr{
		//	Noctty: false,
		//},
	})
	if err != nil {
		log.Panicf("error starting process: %s", err)
	}
	//proc.Release()

	var procState *os.ProcessState
	procState, err = proc.Wait()
	if err != nil {
		log.Panicf("error waiting for script: %s", err)
	}
	log.Print(procState.String())

	err = stdoutW.Close()
	if err != nil {
		log.Panicf("error closing stdout writer: %s", err)
	}

	err = stderrW.Close()
	if err != nil {
		log.Panicf("error closing stderr writer: %s", err)
	}

	err = stdinR.Close()
	if err != nil {
		log.Panicf("error closing stdin reader: %s", err)
	}

	err = conn.Close()
	if err != nil {
		log.Panicf("error closing the connection: %s", err)
	}
}
