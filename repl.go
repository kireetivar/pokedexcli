package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	availableCommands := getCommands()

	for {
		fmt.Print(" > ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0{
			continue
		}

		commandName := cleaned[0]
		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Printf("Invalid command: %v\n",commandName)
			continue
		}
		command.callback()
	}
}

type cliCommand struct {
	name string
	description string
	callback func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Prints the help menu",
			callback: callbackHelp,
		},
		"exit": {
			name: "exit",
			description: "Turn off the Pokedex",
			callback: callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(strings.TrimSpace(str))
	words := strings.Fields(lowered)
	return words
}