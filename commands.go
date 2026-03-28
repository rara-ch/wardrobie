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

func buildCommands() commands {
	commands := commands{}

	commands.addCommand(command{
		name:        "add",
		description: "",
		handler:     addHandler,
	})

	commands.addCommand(command{
		name:        "reset",
		description: "",
		handler:     resetHandler,
	})

	commands.addCommand(command{
		name:        "get",
		description: "",
		handler:     getHandler,
	})

	return commands
}
