package main

import (
	"github.com/localhots/broseph/broseph"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-ch
		os.Exit(1)
	}()

	broseph.StartServer()
}
