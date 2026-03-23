package main

type Handler func() error

type command struct {
	name        string
	description string
	handler     Handler
}

type Commands map[string]command

func (commands Commands) AddCommand(name, desciption string, handler Handler) {
	command := command{
		name:        name,
		description: desciption,
		handler:     handler,
	}
	commands[command.name] = command
}
