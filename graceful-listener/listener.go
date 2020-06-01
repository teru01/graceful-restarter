// package graceful-listener provides listener library interfaces
package listener

import (
	"net"
	"os"
	"sync"
	"fmt"
	"os/signal"
	"syscall"
)

type GracefulListener struct {
	listener net.Listener
	wg sync.WaitGroup
	quit chan struct{}
}

func (l *GracefulListener) WaitShutdownAll() {
	close(l.quit)
	l.listener.Close()
	l.wg.Wait()
}

func (l *GracefulListener) WaitAndGracefulShutdown() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM)
	<- ch
	l.WaitShutdownAll()
}

// handle incomming connection and call server.Handler
func (l *GracefulListener) Serve(handler func(net.Conn)) {
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
				handler(conn)
				l.wg.Done()
				conn.Close()
			}()
		}
	}
}

// Listen use file descriptor passed by Master as a socket.
func Listen() (*GracefulListener, error) {
	ln, err := net.FileListener(os.NewFile(uintptr(3), "dummy"))
	if err != nil {
		return nil, err
	}
	return &GracefulListener {
		listener: ln,
		quit: make(chan struct{}),
	}, nil
}


