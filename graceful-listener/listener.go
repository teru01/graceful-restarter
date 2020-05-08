// package graceful-listener provides listener library interfaces
package listener

import "net"

// Listen use file descriptor passed by Master as a socket.
func Listen() (net.Listener, error) {
	return net.Listen("tcp", "localhost:8080")
}
