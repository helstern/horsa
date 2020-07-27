package main

import (
	"github.com/coreos/go-systemd/v22/activation"
	"log"
	"os"
)

func main() {
	listeners, err := activation.Listeners() // ❶
	if err != nil {
		log.Panicf("cannot retrieve listeners: %s", err)
	}
	if len(listeners) != 1 {
		log.Panicf("unexpected number of socket activation (%d != 1)",
			len(listeners))
	}

	args := os.Args[1:]
	if len(args) < 1 {
		log.Panicf("unexpected number of arguments (%d < 1)",
			len(args))
	}

	listener := listeners[0]
	for {
		conn, _ := listener.Accept()
		go serve(conn, args)
	}
}
