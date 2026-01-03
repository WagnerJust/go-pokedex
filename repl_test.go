package main

import (
	"testing"

	"github.com/WagnerJust/go-pokedex/internal/pokeapi"
	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"simple hello world": {input: "hello world", expected: []string{"hello", "world"}},
		"different cases":    {input: "Charmander Bulbasaur PIKACHU", expected: []string{"charmander", "bulbasaur", "pikachu"}},
		"extra spaces":    {input: "    YoUr      Mom     ", expected: []string{"your", "mom"}},
	}

	for name, tc := range cases {
        t.Run(name, func(t *testing.T) {
            actual := cleanInput(tc.input)
            diff := cmp.Diff(tc.expected, actual)
            if diff != "" {
                t.Fatal(diff)
            }
        })
    }
}

func TestDisplayPokemon(t *testing.T) {
	cases := map[string]struct {
		input    pokeapi.Pokemon
		expected string
	}{
		"Emolga": {
			input: pokeapi.Pokemon{
				Name:   "Emolga",
				Height: 4,
				Weight: 60,
				Stats: []pokeapi.Stats{
					{
						BaseStat: 20,
						Stat: pokeapi.Stat{
							Name: "hp",
						},
					},
					{
						BaseStat: 55,
						Stat: pokeapi.Stat{
							Name: "attack",
						},
					},
				},
				Types: []pokeapi.Types{
					{
						Type: pokeapi.Type{
							Name: "electric",
						},
					},
					{
						Type: pokeapi.Type{
							Name: "flying",
						},
					},
				},
			},
			expected: `Name: Emolga
Height: 4
Weight: 60
Stats:
	-hp: 20
	-attack: 55
Types:
	-electric
	-flying
`,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual, err := displayPokemon(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			diff := cmp.Diff(tc.expected, actual)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
