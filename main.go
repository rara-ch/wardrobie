package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rara-ch/wardrobie/internal/database"
)

type state struct {
	db *database.Queries
}

func main() {
	fmt.Println("Wardrobie begins here")

	fmt.Println("Loading .env file")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("could not load environment variables: %s", err)
	}

	fmt.Println("Connecting to database")
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("could not connect to database: %s", err)
	}
	dbQueries := database.New(db)

	state := &state{
		db: dbQueries,
	}

	fmt.Println("Building commands")
	commands := buildCommands()

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalln("no command given")
	} else {
		if command, ok := commands[args[0]]; ok {
			fmt.Println("Running command")
			err := command.handler(state, args[1:])
			if err != nil {
				log.Fatalf("error when running command %s: %s", args[0], err)
			}
		} else {
			log.Fatalf("command does not exist")
		}
	}
}
