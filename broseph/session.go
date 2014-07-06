package broseph

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/kr/pty"
	"io"
	"net"
	"os/exec"
)

const (
	BufferSize = 1024 * 16
)

var (
	sessions = []Session{}
)

type (
	Session struct {
		Id string
	}
)

func NewSession() *Session {
	s := Session{
		Id: uuid.New(),
	}
	go s.Accept()
	sessions = append(sessions, s)

	return &s
}

func (s *Session) Accept() {
	l, err := net.Listen("unix", "/tmp/"+s.Id+".sock")
	if err != nil {
		panic(err)
	}

	bash := exec.Command("bash")
	_pty, err := pty.Start(bash)
	if err != nil {
		println("Failed starting bash", err.Error())
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			println("Error accepting connection", err.Error())
			continue
		}

		io.Copy(_pty, conn)
		io.Copy(conn, _pty)
	}
}
