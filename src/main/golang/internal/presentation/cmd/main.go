package cmd

import "github.com/spf13/cobra"

var (
	// Used for flags.
	//cfgFile     string
	//userLicense string

	descShort string = "socket server proxy"
	descLong  string = `horsa is a socket server proxy that simplifies interaction with a socket by exposing an incoming connection to a handler process.`
)

type SystemdCommand = cobra.Command
type RootCommand = cobra.Command

func NewRootCommand(systemd *SystemdCommand) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "horsa",
		Short: descShort,
		Long:  descLong,
	}

	//SetFlagsForConfiguration(rootCmd)
	rootCmd.AddCommand(
		systemd,
	)

	return rootCmd
}
