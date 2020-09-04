
package config

import (
    "fmt"
    "strings"

    "github.com/codetojoy/config/json"
)

type PlayerConfig struct {
    Name string
    WhichStrategy string
}

type Config struct {
    NumCards int
    NumGames int
    NumPlayers int
    NumCardsPerHand int
    MaxCard int
    Players []PlayerConfig
}

func NewConfigFromFile(jsonFile string) Config {
    jsonConfig := json.BuildJsonConfig(jsonFile)
    // jsonConfig.Log()
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

func NewConfig() Config {
    const numCards = 12
    const numPlayers = 3
    const numGames = 10
    // numPlayers + 1 because of kitty
    const numCardsPerHand = numCards / (numPlayers + 1)
    const maxCard = numCards

    return Config{NumCards: numCards, NumGames: numGames,
                    NumPlayers: numPlayers, NumCardsPerHand: numCardsPerHand,
                    MaxCard: maxCard, Players: []PlayerConfig{}}
}

func NewConfigForTesting(numCards int, numPlayers int, numGames int) Config {
    // numPlayers + 1 because of kitty
    numCardsPerHand := numCards / (numPlayers + 1)
    maxCard := numCards

    return Config{NumCards: numCards, NumGames: numGames,
                    NumPlayers: numPlayers, NumCardsPerHand: numCardsPerHand,
                    MaxCard: maxCard, Players: []PlayerConfig{}}
}

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
