package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Wardrobie begins here")

	args := os.Args[1:]

	for index, arg := range args {
		fmt.Printf("Arg %d: %s\n", index+1, arg)
	}
}
