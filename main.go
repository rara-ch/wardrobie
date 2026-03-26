package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Wardrobie begins here")

	fmt.Println("Building commands")
	commands := commands{}
	commands.AddCommand("command1", "first test command", func() error {
		fmt.Println("first command run")
		return nil
	})
	commands.AddCommand("command2", "second test command", func() error {
		fmt.Println("second command run")
		return nil
	})

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalln("no command given")
	} else {
		if command, ok := commands[args[0]]; ok {
			command.handler()
		} else {
			log.Fatalf("command does not exist")
		}
	}
}
