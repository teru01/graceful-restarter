package server

import (
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Master struct {
	listener net.Listener
	sigCh chan os.Signal
	workerCh chan WorkerStatus
}

type WorkerStatus struct {
	status int
}

func NewMaster(addr string) (*Master, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Master{
		listener: l,
		sigCh: make(chan os.Signal, 1),
		workerCh: make(chan WorkerStatus, 5), // TODO: define chan size
	}, nil
}

func (master *Master) Run() {
	// create worker and pass socket discriptor, then wait until received signal
	signal.Notify(master.sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	for {
		select {
		case c := <- master.sigCh:
			//
			_ = c
		case c := <- master.workerCh:
			_ = c
		}

	}
}

// CreateWorker creates listener process and return created process struct.
func (m *Master) CreateWorker() os.Process {
	return os.Process{}

	// 死んだら呼び出し元に通知
}
