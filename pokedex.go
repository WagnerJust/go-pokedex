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

func (p *Pokedex) Add(pokemon pokeapi.Pokemon) {
	p.pokemon[pokemon.Name] = pokemon
}

func (p *Pokedex) Get(name string) (pokeapi.Pokemon, bool) {
	if pokemon, ok := p.pokemon[name]; ok {
		return pokemon, true
	}
	return pokeapi.Pokemon{}, false
}
