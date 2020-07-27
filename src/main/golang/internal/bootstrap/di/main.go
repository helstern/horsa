package di

import (
	"context"
	"github.com/sarulabs/di/v2"
)

func NewContainer(ctx context.Context) (di.Container, error) {
	builder, _ := di.NewBuilder()
	var err error

	err = RootCommand.Register(ctx, builder)
	if err != nil {
		return nil, err
	}

	ctn := builder.Build()
	return ctn, nil
}
