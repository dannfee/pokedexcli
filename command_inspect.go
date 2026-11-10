package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, ok := cfg.caughtPokemon[name]
	if ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, s := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, t := range pokemon.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}
		return nil
	}
	fmt.Println("you have not caught that pokemon")
	return nil
}
