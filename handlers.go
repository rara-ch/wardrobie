package main

type handler func() error

type command struct {
	name        string
	description string
	handler     handler
}

type commands map[string]command

func (commands commands) AddCommand(name, desciption string, handler handler) {
	command := command{
		name:        name,
		description: desciption,
		handler:     handler,
	}
	commands[command.name] = command
}
