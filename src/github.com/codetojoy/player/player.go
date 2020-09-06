package player

import (
	"fmt"
	"strings"

	"github.com/codetojoy/strategy"
)

type Player struct {
	name         string
	hand         Hand
	strategy     strategy.Strategy
	PlayerStats  PlayerStats
	offer        int
	channel      chan int
	isOfferReady bool
}

type Metric func(*Player) int

func ByGameTotal(player *Player) int {
	return player.PlayerStats.GameTotal
}

func ByNumGamesWon(player *Player) int {
	return player.PlayerStats.NumGamesWon
}

func ByOffer(player *Player) int {
	return player.offer
}

func NewPlayer(name string, whichStrategy string) Player {
	ch := make(chan int)
	return Player{name: name, strategy: strategy.BuildStrategy(whichStrategy),
		hand: NewHandNoCards(), PlayerStats: PlayerStats{}, channel: ch}
}

func (player *Player) SetHand(hand Hand) {
	player.hand = hand
}

func (player *Player) GetName() string {
	return player.name
}

// this will block on channel
func (player *Player) GetOffer() int {
	result := 0

	if player.getIsOfferReady() {
		result = player.offer
		// in event of being called repeatedly, this is safe. i.e. it can handle if target is not present
		player.hand.RemoveCard(player.offer)
	}

	return result
}

// this will block on channel
func (player *Player) getIsOfferReady() bool {
	// result is always true, since we are blocking until it is true
	result := true

	if !player.isOfferReady {
		player.offer = <-player.channel
		player.isOfferReady = true
	}

	return result
}

func (player *Player) MakeOffer(prizeCard int, maxCard int) {
	player.isOfferReady = false
	go player.strategy.SelectCard(player.channel, prizeCard, player.hand.GetCards(), maxCard)
	// offer := player.strategy.SelectCard(prizeCard, player.hand.GetCards(), maxCard)
	// player.offer = offer
	// player.hand.RemoveCard(offer)
}

func (player *Player) WinsRound(prizeCard int) {
	player.PlayerStats.GameTotal += prizeCard
	player.PlayerStats.NumRoundsWon += 1
}

func (player *Player) WinsGame() {
	player.PlayerStats.NumGamesWon += 1
}

func (player *Player) ClearGameStats() {
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
