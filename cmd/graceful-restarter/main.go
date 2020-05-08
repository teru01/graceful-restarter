package main

import (
	"flag"
	"fmt"
)

func main() {
	var addr = flag.String("L", "127.0.0.1:0", "listen addr:port")
	fmt.Println(*addr)
}
