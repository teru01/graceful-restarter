package main

import (
	listener "github.com/teru01/graceful-restarter/graceful-listener"
)

/*
This is an example using listener
*/

func main() {
	l, err := listener.Listen()
	_ = l
	_ = err
}
