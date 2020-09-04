// Package config handles configuration of the game and players.
package config

import (
	"fmt"
	"strings"

	"github.com/codetojoy/config/json"
)

type PlayerConfig struct {
	Name          string
	WhichStrategy string
}

type Config struct {
	NumCards        int
	NumGames        int
	NumPlayers      int
	NumCardsPerHand int
	MaxCard         int
	Players         []PlayerConfig
}

// NewConfigFromFile builds a Config from a JSON file.
func NewConfigFromFile(jsonFile string) Config {
	jsonConfig := json.BuildJsonConfig(jsonFile)
	numCards := jsonConfig.NumCards
	numGames := jsonConfig.NumGames
	numPlayers := len(jsonConfig.Players)
	config := NewConfigForTesting(numCards, numPlayers, numGames)

	for _, jsonPlayer := range jsonConfig.Players {
		name := jsonPlayer.Name
		whichStrategy := jsonPlayer.Strategy
		playerConfig := PlayerConfig{Name: name, WhichStrategy: whichStrategy}
		config.Players = append(config.Players, playerConfig)
	}

	return config
}

// NewConfigForTesting builds a Config in a unit-test context.
func NewConfigForTesting(numCards int, numPlayers int, numGames int) Config {
	// numPlayers + 1 because of kitty
	numCardsPerHand := numCards / (numPlayers + 1)
	maxCard := numCards

	return Config{NumCards: numCards, NumGames: numGames,
		NumPlayers: numPlayers, NumCardsPerHand: numCardsPerHand,
		MaxCard: maxCard, Players: []PlayerConfig{}}
}

// String is standard functionality.
func (config *Config) String() string {
	result := strings.Builder{}

	result.WriteString(fmt.Sprintf("numCards: %d numGames: %d numPlayers: %d\n", config.NumCards,
		config.NumGames, config.NumPlayers))

	for _, playerConfig := range config.Players {
		name := playerConfig.Name
		whichStrategy := playerConfig.WhichStrategy
		result.WriteString(fmt.Sprintf("name: %v strategy: %v\n", name, whichStrategy))
	}

	return result.String()
}
