package broseph

import (
	"github.com/kr/pty"
)

var Sessions []Session

type Session struct {
	ID string
}

func NewSession() Session {
	return Session{}
}
