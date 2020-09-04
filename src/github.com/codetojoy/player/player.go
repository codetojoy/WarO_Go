
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
    PlayerStats PlayerStats
}

func NewPlayer(name string, whichStrategy string) Player {
    return Player{name: name, strategy: strategy.BuildStrategy(whichStrategy),
                    hand: NewHandNoCards(), PlayerStats: PlayerStats{}}
}

func BuildPlayers(config config.Config) []Player {
    result := []Player{}

    for _, configPlayer := range config.Players {
        name := configPlayer.Name
        whichStrategy := configPlayer.WhichStrategy
        player := NewPlayer(name, whichStrategy)
        result = append(result, player)
    }

    return result
}

func (player *Player) SetHand(hand Hand) {
    player.hand = hand
}

func (player *Player) GetName() string {
    return player.name
}

func (player *Player) LogCards(prefix string) {
    fmt.Printf("TRACER logCards %v %v cards: %v\n", prefix, player.name, player.hand.GetCards())
}

func LogCardsForPlayers(players []Player, prefix string) {
    for _, player := range players {
        player.LogCards(prefix)
    }
}

func (player *Player) GetOffer(prizeCard int, maxCard int) int {
    offer := player.strategy.SelectCard(prizeCard, player.hand.GetCards(), maxCard)
    player.hand.RemoveCard(offer)
    return offer
}

func (player *Player) WinsRound(prizeCard int) {
    player.PlayerStats.GameTotal += prizeCard
    player.PlayerStats.NumRoundsWon += 1
}

func (player *Player) WinsGame() {
    player.PlayerStats.NumGamesWon += 1
}

func (player *Player) NewGame() {
    player.PlayerStats.GameTotal = 0
    player.PlayerStats.NumRoundsWon = 0
}

func (player *Player) String() string {
    result := strings.Builder{}

    result.WriteString(fmt.Sprintf("%s : ", player.name))
    result.WriteString(fmt.Sprintf("%s ", player.PlayerStats.String()))
    result.WriteString(fmt.Sprintf("%s ", player.hand.String()))

    return result.String()
}

// -------- for test

func BuildPlayerForTesting(name string, whichStrategy string, cards []int) Player {
    player := NewPlayer(name, whichStrategy)
    player.SetHand(NewHand(cards))
    return player
}

func (player *Player) GetCardsForTesting() []int {
    return player.hand.GetCards()
}
