package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"text/template"
	"github.com/WagnerJust/go-pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient *pokeapi.PokeApiClient
	nextUrl		  *string
	previousUrl   *string
	pokedex 	  *Pokedex
}
type cliCommand struct {
	name, description string
	callback          func(config *config, args ...string) error
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
		"explore": {
			name:        "explore",
			description: "Explore a location area",
			callback: exploreLocationArea,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    catchPokemon,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemon",
			callback:    inspectPokemon,
		},
	}
}

func commandExit(config *config, args ...string ) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func help(config *config, args ...string ) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, commandInfo := range commands {
		fmt.Printf("%s: %s\n", name, commandInfo.description)
	}
	return nil
}

func mapLocationAreas(config *config, args ...string ) error {
	locationAreas, err := config.pokeapiClient.GetLocationAreas(config.nextUrl)
	if err != nil {
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

func mapLocationAreasB(config *config, args ...string ) error {
	locationAreas, err := config.pokeapiClient.GetLocationAreas(config.previousUrl)
	if err != nil {
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

func exploreLocationArea(config *config, args ...string ) error {
	if len(args) < 2 {
		return fmt.Errorf("explore requires a location area name")
	}
	name := args[1]
	fmt.Printf("Exploring %s...\n", name )
	detailedLocationArea, err := config.pokeapiClient.GetDetailedLocationArea(name)
	if err != nil {
		return err
	}
	if len(detailedLocationArea.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon in this area")
		return nil
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range detailedLocationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func catchPokemon(config *config, args ...string ) error {
	if len(args) < 2 {
		return fmt.Errorf("catch requires a pokemon name")
	}
	name := args[1]
	fmt.Printf("Throwing a Pokeball at %s...\n", name )
	pokemon, err := config.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	if !attemptToCatch(pokemon) {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	config.pokedex.Add(pokemon)
	return nil
}

func displayPokemon(pokemon pokeapi.Pokemon) (string, error) {
	pokemonInfo := `Name: {{.Name}}
Height: {{.Height}}
Weight: {{.Weight}}
Stats:
{{- range .Stats}}
	-{{.Stat.Name}}: {{.BaseStat}}
{{- end}}
Types:
{{- range .Types}}
	-{{.Type.Name}}
{{- end}}
`

	pokemonTemplate, err := template.New("pokemonInfo").Parse(pokemonInfo)
	if err != nil {
		return "", err
	}
	var result strings.Builder
	err = pokemonTemplate.Execute(&result, pokemon)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}
func inspectPokemon(config *config, args ...string) error {
	if len(args) < 2 {
		return fmt.Errorf("inspect requires a pokemon name")
	}
	name := args[1]

	pokemon, found := config.pokedex.Get(name)
	if !found {
		return fmt.Errorf("you can only inspect caught pokemon. You have not caught a %s", name)
	}
	output, err := displayPokemon(pokemon)
	if err != nil {
		return err
	}
	fmt.Print(output)
	return nil


}
func cleanInput(text string) []string {
	fields := strings.Fields(text)
	for index, word := range fields {
		fields[index] = strings.ToLower(word)
	}
	return fields
}

func attemptToCatch(pokemon pokeapi.Pokemon) bool {
	result := rand.Intn(pokemon.BaseExperience)
	if result % 4 == 0 {
		return true
	}
	return false
}
func ReplLoop() {
	config := config{}
	config.pokeapiClient = pokeapi.NewPokeApiClient()
	config.pokedex = NewPokedex()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		userCommand := input[0]
		value, ok := commands[userCommand]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := value.callback(&config, input...)
		if err != nil {
			fmt.Println(err)
		}
	}
}
