package broseph

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/kr/pty"
)

var Sessions map[UUID]Session

type (
	UUID    string
	Session struct {
		ID UUID
	}
)

func NewSession() Session {
	s := Session{
		ID: uuid.New(),
	}
}
