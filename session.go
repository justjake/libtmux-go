package tmux

import (
	"regexp"
	"strings"
)

const None string = ""

type Session struct {
	Name string
}

type DisplayOptions struct {
	Background string
	Foreground string
	Time       string
}

func (s *Session) Set(key, value string) {
	if value == "" {
		Run("set", "-q", "-u", "-t", s.Name, key)
	} else {
		Run("set", "-q", "-t", s.Name, key, value)
	}
}

func NewDisplayOptions() *DisplayOptions {
	return &DisplayOptions{
		Background: Yellow,
		Foreground: Black,
		Time:       "5000",
	}
}

func (s *Session) ApplyDisplayOptions(option *DisplayOptions) {
	styleBuilder := []string{}

	if option.Foreground != "" {
		styleBuilder = append(styleBuilder, "fg="+option.Foreground)
	}

	if option.Background != "" {
		styleBuilder = append(styleBuilder, "bg="+option.Background)
	}
	style := strings.Join(styleBuilder, ",")
	s.Set("message-style", style)
	s.Set("display-time", option.Time)
}

func (s *Session) GetDisplayOptions() *DisplayOptions {
	d := &DisplayOptions{}

	message_style := Run("show", "-t", s.Name, "message-style")
	message_style = strings.TrimSpace(message_style)
	re := regexp.MustCompile("^(?:message-style (?:fg=(.*?))?,?(?:bg=(.*?))?)?$")
	styles := re.FindStringSubmatch(message_style)

	d.Foreground = styles[1]
	d.Background = styles[2]
	d.Time = Run("show", "-t", s.Name, "display-time")
	return d
}
