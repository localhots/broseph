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
	// Session is a container responsible for storing session-related
	// information.
	Session struct {
		Id string // Session identifier
	}
)

// Creates a new Session and starts accepting connections.
// Returns a pointer to a newly created session.
func NewSession() *Session {
	s := Session{
		Id: uuid.New(),
	}
	go s.Accept()
	sessions = append(sessions, s)

	return &s
}

// Starts accepting connections on the UNIX domain socket associated with
// session.
func (s *Session) Accept() {
	l, err := net.Listen("unix", "/tmp/"+s.Id+".sock")
	if err != nil {
		println("Failed opening socket:", err.Error())
		return
	}

	bash := exec.Command("bash")
	_pty, err := pty.Start(bash)
	if err != nil {
		println("Failed starting bash:", err.Error())
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			println("Error accepting connection:", err.Error())
			continue
		}

		// Half-duplex
		// TODO: Full-duplex
		io.Copy(_pty, conn)
		io.Copy(conn, _pty)
	}
}
