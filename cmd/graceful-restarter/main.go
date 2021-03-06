// package main defines command.
// example: ./graceful-restarter -L 127.0.0.1:8080 [server-command] [args...]
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/teru01/graceful-restarter"
)

func help() {
	fmt.Fprintf(os.Stderr, "Usage: graceful-restarter -L [addr:port] program arg-1 arg-2 ...\n")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "specify executable path")
		help()
		os.Exit(1)
	}
	addr := flag.String("L", "127.0.0.1:0", "listen addr:port")
	flag.Parse()
	fmt.Println(*addr)
	master, err := server.NewMaster(*addr, flag.Args())
	if err != nil {
		panic(err)
	}
	if err := master.Run(); err != nil {
		panic(err)
	}
}
