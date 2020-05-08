package server

import (
	"net"
	"os"
)

type Master struct {
	listener net.Listener
}

func NewMaster(addr net.Addr) (*Master, error) {
// create master
	l, err := net.Listen(addr.Network(), addr.String())
	if err != nil {
		return nil, err
	}
	return &Master{
		listener: l,
	}
}

func (m *Master) Run() {
// create worker and pass socket discriptor, then wait until received signal
	
}

// CreateWorker creates listener process and return created process struct.
func (m *Master) CreateWorker() os.Process {

}
