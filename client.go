package tmux

import "strings"

// path to character device file
type Client struct {
	TTY string
}

func (c Client) String() string {
	return c.TTY
}

func GetAllClients() []Client {
	s := Run("list-clients", "-F", "#{client_tty}")
	ttys := strings.Split(strings.TrimSpace(s), "\n")
	clients := make([]Client, len(ttys))

	for i := range ttys {
		clients[i] = Client{ttys[i]}
	}

	return clients
}
