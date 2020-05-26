// package graceful-listener provides listener library interfaces
package listener

import (
	"net"
	"os"
	"sync"
	"os/signal"
	"syscall"
)

type GracefulListener {
	listener net.Listener
	wg sync.WaitGroup
	quit chan struct{}
}

func (l *GracefulListener) WaitShutdownAll() {
	close(quit)
	wg.Wait()
}

// func (l *GracefulListener) Accept() (net.Conn, error) {
// 	l.wg.Add()
// }

// handle incomming connection and call server.Handler
func (l *GracefulListener) serve(server http.Server) {
	for {
		conn, err := l.listener.Accept()
		if err != nil {
			
		}
	}
}

// Listen use file descriptor passed by Master as a socket.
func Listen() (net.Listener, error) {
	return GraceGracefulListener {
		listener: net.FileListener(os.NewFile(uintptr(3), "dummy")),
		quit: make(chan struct{})
	}
}

func WaitAndGracefulShutdown(l *net.Listener) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM)
	<- ch
	l.WaitShutdownAll()
}

