package main

// type handler func() error

type command struct {
	name        string
	description string
	handler     func(s *state, args []string) error
}

type commands map[string]command

func (commands commands) addCommand(command command) {
	commands[command.name] = command
}
