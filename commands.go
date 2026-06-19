package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.commands[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, handler func(*state, command) error) {
	if c.commands == nil {
		c.commands = make(map[string]func(*state, command) error)
	}
	c.commands[name] = handler
}
