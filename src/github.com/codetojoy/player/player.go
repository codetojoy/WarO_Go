
package player

import (
    "fmt"
    "strings"

    "github.com/codetojoy/config"
    "github.com/codetojoy/strategy"
)

type Player struct {
    name string
    hand Hand
    strategy strategy.Strategy
    gameTotal int
    numRoundsWon int
    numGamesWon int
}

func NewPlayer(name string) Player {
    return Player{name: name, strategy: strategy.Strategy{}}
}

func BuildPlayers(config config.Config) []Player {
    result := []Player{}

    for i := 1; i <= config.NumPlayers; i++ {
        name := fmt.Sprintf("player%d", i)
        player := NewPlayer(name)
        result = append(result, player)
    }

    return result
}

func (player *Player) SetHand(hand Hand) {
    player.hand = hand
}

func (player *Player) GetCardsForTesting() []int {
    return player.hand.GetCards()
}

func (player *Player) GetName() string {
    return player.name
}

func (player *Player) GetOffer(prizeCard int, maxCard int) int {
    offer := player.strategy.SelectCard(prizeCard, player.hand.GetCards(), maxCard)
    player.hand.RemoveCard(offer)
    return offer
}

func (player *Player) String() string {
    result := strings.Builder{}

    result.WriteString(fmt.Sprintf("%s : ", player.name))
    result.WriteString(fmt.Sprintf("%s ", player.hand.String()))

    return result.String()
}
