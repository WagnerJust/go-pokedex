package main

import (
	"testing"
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
