package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"

	listener "github.com/teru01/graceful-restarter/graceful-listener"
)

func handler(conn net.Conn) {
	fmt.Fprintf(conn, "server pid %d\n", os.Getpid())
	r := bufio.NewReader(conn)
	for {
		content, err := r.ReadString('\n')
		if err == io.EOF {
			return
		} else if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
		fmt.Fprintf(conn, "pid: %d, %v", os.Getpid(), content)
	}
}

func main() {
	l, err := listener.Listen()
	if err != nil {
		panic(err)
	}
	go l.Serve(handler)
	l.WaitAndGracefulShutdown()
}
