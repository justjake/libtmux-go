// Package tmux provides primitives for interactive with tmux
package tmux

import (
	"log"
	"os/exec"
)

const (
	baseCommand = "tmux"
)

func Display(message string, clients ...Client) {
	for i := range clients {
		Run("display", "-c", clients[i].TTY, message)
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
