package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/WagnerJust/go-pokedex/internal/pokeapi"
)
type config struct {
	pokeapiClient *pokeapi.PokeApiClient
	nextUrl		  *string
	previousUrl   *string
}

type cliCommand struct {
	name, description string
	callback          func(config *config) error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    help,
		},
		"map": {
			name:	"map",
			description: "Shows you the next 20 location areas",
			callback: mapLocationAreas,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows you the previous 20 location areas",
			callback:    mapLocationAreasB,
		},
	}
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func help(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, commandInfo := range commands {
		fmt.Printf("%s: %s\n", name, commandInfo.description)
	}
	return nil
}

func mapLocationAreas(config *config) error {
	locationAreas, err := config.pokeapiClient.GetLocationAreas(config.nextUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	config.nextUrl = locationAreas.Next
	config.previousUrl = locationAreas.Previous
	if config.nextUrl == nil {
		fmt.Println("you're on the last page")
	}
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func mapLocationAreasB(config *config) error {
	locationAreas, err := config.pokeapiClient.GetLocationAreas(config.previousUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	config.previousUrl = locationAreas.Previous
	config.nextUrl = locationAreas.Next
	if config.previousUrl == nil {
		fmt.Println("you're on the first page")
	}
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	return nil
}


func cleanInput(text string) []string {
	fields := strings.Fields(text)
	for index, word := range fields {
		fields[index] = strings.ToLower(word)
	}
	return fields
}

func ReplLoop() {
	config := config{}
	config.pokeapiClient = pokeapi.NewPokeApiClient()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		userCommand := input[0]
		value, ok := commands[userCommand]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		value.callback(&config)
	}
}
