package main

import (
	"context"
	"github.com/helstern/horsa/src/main/golang/internal/bootstrap/di"
	"log"
)

func main() {
	ctx := context.Background()
	container, _ := di.NewContainer(ctx)

	command := di.RootCommand.Get(container)
	err := command.ExecuteContext(ctx)

	if err != nil {
		log.Fatalf("failed to run %s", err)
	}
}
