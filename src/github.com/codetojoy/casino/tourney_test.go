package casino

import (
	"testing"

	"github.com/codetojoy/config"
	"github.com/codetojoy/player"
	"github.com/codetojoy/strategy"
)

func TestPlayTourneyWinnerCase1(t *testing.T) {
	p1 := player.NewPlayer("beethoven", strategy.NEXT_CARD, "")
	p2 := player.NewPlayer("chopin", strategy.NEXT_CARD, "")
	p3 := player.NewPlayer("mozart", strategy.NEXT_CARD, "")

	players := []player.Player{p1, p2, p3}

	const numCards = 12
	numPlayers := len(players)
	const numGames = 10
	config := config.NewConfigForTesting(numCards, numPlayers, numGames)

	deckProvider := SimpleDeckProvider{}
	dealer := NewTestDealer(deckProvider)

	// kitty: 1,2,3
	// beethoven: 4,5,6
	// chopin: 7,8,9
	// mozart: 10,11,12

	// test
	PlayTourney(config, players, dealer)
	result := players[2].PlayerStats.NumGamesWon

	expected := numGames
	ok := (result == expected)

	if !ok {
		t.Errorf("PlayTourney() error actual: %d expected: %d", result, expected)
	}
}

func TestPlayTourneyWinnerCase2(t *testing.T) {
	p1 := player.NewPlayer("beethoven", strategy.NEXT_CARD, "")
	p2 := player.NewPlayer("chopin", strategy.NEXT_CARD, "")
	p3 := player.NewPlayer("mozart", strategy.NEXT_CARD, "")

	players := []player.Player{p1, p2, p3}

	const numCards = 12
	numPlayers := len(players)
	const numGames = 10
	config := config.NewConfigForTesting(numCards, numPlayers, numGames)

	// kitty: 2,3,1
	// beethoven: 4,11,6
	// chopin: 7,8,12
	// mozart: 10,5,9

	deck := []int{2, 3, 1, 4, 11, 6, 7, 8, 12, 10, 5, 9}
	deckProvider := TestDeckProvider{deck: deck}
	dealer := NewTestDealer(deckProvider)

	// test
	PlayTourney(config, players, dealer)
	result := players[0].PlayerStats.NumGamesWon

	expected := numGames
	ok := (result == expected)

	if !ok {
		t.Errorf("PlayTourney() error actual: %d expected: %d", result, expected)
	}
}

func TestDetermineTourneyWinner(t *testing.T) {
	p1 := player.NewPlayer("beethoven", strategy.NEXT_CARD, "")
	p2 := player.NewPlayer("chopin", strategy.NEXT_CARD, "")
	p3 := player.NewPlayer("mozart", strategy.NEXT_CARD, "")

	p1.PlayerStats.NumGamesWon = 0
	p2.PlayerStats.NumGamesWon = 3
	p3.PlayerStats.NumGamesWon = 4

	players := []player.Player{p1, p2, p3}

	// test
	result := determineTourneyWinner(players)

	// TODO: figure out how to use test-assert library (see README.md)
	ok := result.GetName() == p3.GetName()

	if !ok {
		t.Errorf("determineTourneyWinner() error actual: %v expected %v", result.GetName(), p3.GetName())
	}
}
