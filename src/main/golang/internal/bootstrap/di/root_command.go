package di

import (
	"context"
	"github.com/helstern/horsa/src/main/golang/internal/presentation/cmd"
	"github.com/helstern/horsa/src/main/golang/internal/presentation/cmd/systemd"
	"github.com/sarulabs/di/v2"
	"reflect"
)

type ProviderRootCommand string

func (this ProviderRootCommand) Get(ctn di.Container) *cmd.RootCommand {
	return ctn.Get(string(this)).(*cmd.RootCommand)
}

func (this ProviderRootCommand) Register(ctx context.Context, builder *di.Builder) error {
	def := di.Def{
		Build: NewRootCommand,
		Name:  string(this),
	}
	return builder.Add(def)
}

var RootCommand = ProviderRootCommand(reflect.TypeOf((*ProviderRootCommand)(nil)).Elem().String())

func NewRootCommand(ctn di.Container) (interface{}, error) {
	systemdCommand := systemd.NewCommand()

	rootCommand := cmd.NewRootCommand(systemdCommand)
	return rootCommand, nil
}
