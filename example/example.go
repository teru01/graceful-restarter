package main

import (
	"fmt"
	"net/http"
	"os"

	listener "github.com/teru01/graceful-restarter/graceful-listener"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "server pid %d\n", os.Getpid())
}

func main() {
	l, err := listener.Listen()
	if err != nil {
		panic(err)
	}
	server := http.Server{
		Handler: http.HandlerFunc(handler),
	}
	go l.serve(server)
	// go server.Serve(l)
	listener.WaitAndGracefulShutdown(l)
}
