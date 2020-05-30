package main

import (
	"fmt"
	"net"
	"os"

	listener "github.com/teru01/graceful-restarter/graceful-listener"
)

func handler(conn net.Conn) {
	fmt.Fprintf(conn, "server pid %d\n", os.Getpid())
}

func main() {
	l, err := listener.Listen()
	if err != nil {
		panic(err)
	}
	go l.Serve(handler)
	// go server.Serve(l)
	l.WaitAndGracefulShutdown()
}
