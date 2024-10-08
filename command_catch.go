package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	fmt.Printf("Throwing a Pokeball at %s\n", pokemon.Name)
	randNum := rand.IntN(pokemon.BaseExperience)

	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught\n", pokemonName)
	fmt.Println("You may now inspect it with the inspect command.")

	return nil
}
