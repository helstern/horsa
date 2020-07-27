package systemd

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/go-systemd/v22/activation"
	"github.com/helstern/horsa/src/main/golang/internal/core/process"
	"github.com/helstern/horsa/src/main/golang/internal/presentation/cmd"
	"github.com/spf13/cobra"
	"log"
	"time"
)

func NewCommand() *cmd.SystemdCommand {
	var command = &cobra.Command{
		Use:   "systemd",
		Short: "proxy a socket managed by systemd to a specified command",
		Long:  `this command tells horsa to listen to incoming connections on the socket, accept them and then fork the command which will have its standard file descriptors bound to the connection`,
		RunE:  Handler,
	}

	return command
}

func Handler(command *cobra.Command, args []string) error {

	listeners, err := activation.Listeners() // ❶
	if err != nil {
		return fmt.Errorf("cannot retrieve listeners: %w", err)
	}
	if len(listeners) != 1 {
		return errors.New(
			fmt.Sprintf("unexpected number of socket activation (%d != 1)", len(listeners)),
		)
	}

	if len(args) < 1 {
		return errors.New(
			fmt.Sprintf("unexpected number of arguments (%d < 1)", len(args)),
		)
	}

	listener := listeners[0]
	for {
		conn, _ := listener.Accept()
		ctx, _ := context.WithTimeout(command.Context(), 5*time.Second)
		go func() {
			err = process.HandleConnection(ctx, conn, args)
			if err != nil {
				log.Print(fmt.Sprintf("error handling connection: %s", err))
			}
		}()
	}
}
