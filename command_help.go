package main

import (
	"fmt"
)

func callbackHelp() error {
	fmt.Println("Welcome to Pokedex help Menu")
	fmt.Println("Here are your available commands")


	availableCommands := getCommands()
	for _,v := range availableCommands {
		fmt.Printf(" - %v	%v\n", v.name,v.description)
	}
	fmt.Println("")
	return nil
}