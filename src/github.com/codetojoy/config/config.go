
package config

import (
    "fmt"
)

type Config struct {
    NumCards int
    NumGames int
    NumPlayers int
    NumCardsPerHand int
    MaxCard int
}

func NewConfig() Config {
    const numCards = 12
    const numPlayers = 3
    const numGames = 1
    // numPlayers + 1 because of kitty
    const numCardsPerHand = numCards / (numPlayers + 1)
    const maxCard = numCards

    return Config{NumCards: numCards, NumGames: numGames,
                    NumPlayers: numPlayers, NumCardsPerHand: numCardsPerHand,
                    MaxCard: maxCard}
}

func NewConfigForTesting(numCards int, numPlayers int, numGames int) Config {
    // numPlayers + 1 because of kitty
    numCardsPerHand := numCards / (numPlayers + 1)
    maxCard := numCards

    return Config{NumCards: numCards, NumGames: numGames,
                    NumPlayers: numPlayers, NumCardsPerHand: numCardsPerHand,
                    MaxCard: maxCard}
}

func (config *Config) String() string {
    return fmt.Sprintf("numCards: %d numGames: %d numPlayers: %d", config.NumCards,
            config.NumGames, config.NumPlayers)
}
