package tmux

import "strings"

// path to character device file
type Client struct {
	CurrentSession *Session
	TTY            string
}

func (c *Client) String() string {
	return c.TTY
}

func GetAllClients() []*Client {
	out := Run("list-clients", "-F", "#{client_tty}:#{client_session}")
	lines := strings.Split(strings.TrimSpace(out), "\n")
	clients := make([]*Client, len(lines))

	for i := range lines {
		info := strings.Split(lines[i], ":")
		clients[i] = &Client{
			TTY:            info[0],
			CurrentSession: &Session{info[1]},
		}
	}

	return clients
}

func (c *Client) Display(message string) {
	Run("display", "-c", c.TTY, message)
}
