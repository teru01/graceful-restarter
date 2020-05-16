// package graceful-listener provides listener library interfaces
package listener

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// Listen use file descriptor passed by Master as a socket.
func Listen() (net.Listener, error) {
	return net.FileListener(os.NewFile(uintptr(3), "dummy"))
}

func WaitAndGracefulShutdown(l net.Listener) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM)
	<- ch
	l.Close() // TODO gracefully
}
