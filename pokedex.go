package main

import (
	"github.com/WagnerJust/go-pokedex/internal/pokeapi"
)
type Pokedex struct {
	pokemon map[string]pokeapi.Pokemon
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		pokemon: make(map[string]pokeapi.Pokemon),
	}
}

func (p *Pokedex) Add(name string, pokemon pokeapi.Pokemon) {
	p.pokemon[name] = pokemon
}

func (p *Pokedex) List() []pokeapi.Pokemon {
	// ...
}
