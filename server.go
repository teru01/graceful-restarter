package server

import (
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

type Master struct {
	listener net.Listener
	command string
	commandArgs []string
	sigCh chan os.Signal
	workerCh chan WorkerStatus
}

type WorkerStatus struct {
	exitStatus int
	err error
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
		workerCh := master.CreateWorker()
		select {
		case c := <- master.sigCh:
			//
			_ = c
		case c := <- workerCh:
			// worker exited.
			_ = c
		}

	}
}

// CreateWorker creates listener process and return created process struct.
func (master *Master) CreateWorker() chan WorkerStatus {
	result := make(chan WorkerStatus)

	go func() {
		cmd := exec.Command(master.command, master.commandArgs...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		socketFile, err := master.listener.(*net.TCPListener).File()
		if err != nil {
			result <- WorkerStatus{err: err,}
		}
		cmd.ExtraFiles = append(cmd.ExtraFiles, socketFile)
		err = cmd.Run()
		
	}()

	return result
}
