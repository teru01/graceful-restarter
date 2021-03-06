package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

const FAILSTATUS = 1

func init() {
	log.SetFlags(log.Ltime|log.Lmicroseconds)
}

type Master struct {
	listener net.Listener
	command []string
	sigCh chan os.Signal
	workerCh chan WorkerStatus
}

type WorkerStatus struct {
	pid int
	exitStatus int
	err error
}

func NewMaster(addr string, command []string) (*Master, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	m := &Master{
		listener: l,
		command: command,
		sigCh: make(chan os.Signal, 1),
		workerCh: make(chan WorkerStatus, 5), // TODO: define chan size
	}
	signal.Notify(m.sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	fmt.Printf("master process initiated. pid=%d\n", os.Getpid())
	return m, nil
}

func (master *Master) Run() error {
	// create worker and pass socket discriptor, then wait until received signal
	pid, err := master.CreateWorker()
	if err != nil {
		return err
	}
	for {
		select {
		case c := <- master.sigCh:
			switch c {
			case syscall.SIGHUP:
				newPid, err := master.CreateWorker()
				if err != nil {
					return err
				}
				time.Sleep(1 * time.Second)

				err = killProcess(pid)
				if err != nil {
					return err
				}
				pid = newPid
			case syscall.SIGTERM:
				fallthrough
			case syscall.SIGINT:
				fallthrough
			case syscall.SIGQUIT:
				log.Printf("%v received\n", c)
				err := killProcess(pid)
				return err
			}
		case c := <- master.workerCh:
			// worker exited.
			log.Printf("worker %d exited with status code %d, err %v\n", c.pid, c.exitStatus, c.err)
			if c.err != nil {
				break
			}
		}
	}
}

func killProcess(pid int) error {
	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return p.Signal(syscall.SIGTERM)
}

// CreateWorker creates listener process and return created process struct.
func (master *Master) CreateWorker() (int, error) {
	var cmd *exec.Cmd
	if len(master.command) == 1 {
		cmd = exec.Command(master.command[0])
	} else if len(master.command) > 1 {
		cmd = exec.Command(master.command[0], master.command[1:]...)
	} else {
		return -1, fmt.Errorf("invalid command")
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	socketFile, err := master.listener.(*net.TCPListener).File()
	if err != nil {
		return 0, err
	}
	cmd.ExtraFiles = append(cmd.ExtraFiles, socketFile)
	err = cmd.Start()
	if err != nil {
		return 0, err
	}
	go func() {
		err := cmd.Wait()
		master.workerCh <- WorkerStatus{
			exitStatus: cmd.ProcessState.ExitCode(),
			err: err,
		}
	}()
	return cmd.Process.Pid, nil
}

// func signalToName(s os.Signal) string {

// }
