// package main defines command.
// example: ./graceful-restarter -L 127.0.0.1:8080 [server-command] [args...]
package main

import (
	"flag"
	"fmt"

	"github.com/teru01/graceful-restarter"
)

func main() {
	addr := flag.String("L", "127.0.0.1:0", "listen addr:port")
	fmt.Println(*addr)
	s, _ := server.NewMaster(*addr)
	_ = s
}
