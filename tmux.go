// Package tmux provides primitives for interactive with tmux
package tmux

import (
	"log"
	"os/exec"
	"strconv"
	"time"
)

const (
	baseCommand = "tmux"
)

func Display(message string, options *DisplayOptions, clients ...*Client) {
	storedOptions := make([]*DisplayOptions, len(clients))
	for i := range clients {
		session := clients[i].CurrentSession
		storedOptions[i] = session.GetDisplayOptions()
		session.ApplyDisplayOptions(options)
		clients[i].Display(message)
	}
	milli, err := strconv.Atoi(options.Time)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Duration(milli) * time.Millisecond)

	for i := range clients {
		session := clients[i].CurrentSession
		session.ApplyDisplayOptions(storedOptions[i])
	}
}

func Run(args ...string) string {
	cmd := exec.Command(baseCommand, args...)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	s := string(out[:])
	return s
}
