package main

import (
	"github.com/teru01/graceful-restarter/listener"
)

/*
This is an example using listener
*/

func main() {
	l, err := listener.Listen()
	_ = l
	_ = err
}
