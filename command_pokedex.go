package main

import (
	"fmt"
)

func commandPokedex(cfg *config, _ ...string) error {

	pokedex := cfg.caughtPokemon

	if len(pokedex) > 0 {
		fmt.Printf("Your Pokedex:\n")
		for _, pokemon := range pokedex {
			fmt.Printf(" - %s\n", pokemon.Name)
		}
		return nil
	}

	fmt.Println("you have not caught any pokemon")
	return nil
}
