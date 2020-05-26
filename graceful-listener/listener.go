// package graceful-listener provides listener library interfaces
package listener

import (
	"net"
	"os"
	"sync"
	"os/signal"
	"syscall"
)

type GracefulListener struct {
	listener net.Listener
	wg sync.WaitGroup
	quit chan struct{}
}

func (l *GracefulListener) WaitShutdownAll() {
	close(quit)
	l.listener.Close()
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
			select {
			case <- l.quit:
				return
			default:
				fmt.Fprintln(os.Stderr, err)
			}
		} else {
			l.wg.Add(1)
			go func() {
				server.Handler()
				l.wg.Done()
			}
		}
	}
}

// Listen use file descriptor passed by Master as a socket.
func Listen() (GracefulListener, error) {
	ln, err := net.FileListener(os.NewFile(uintptr(3), "dummy"))
	if err != nil {
		return err
	}
	return GracefulListener {
		listener: ln,
		quit: make(chan struct{})
	}
}

func WaitAndGracefulShutdown(l *net.Listener) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM)
	<- ch
	l.WaitShutdownAll()
}

