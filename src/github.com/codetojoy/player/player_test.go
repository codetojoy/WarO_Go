package player

import (
	"testing"

	"github.com/codetojoy/strategy"
)

func TestWinsGame(t *testing.T) {
	player := NewPlayer("mozart", strategy.NEXT_CARD)

	// test
	player.WinsGame()

	ok := (player.PlayerStats.NumGamesWon == 1)

	if !ok {
		t.Errorf("WinsGame error")
	}
}

func TestWinsRound(t *testing.T) {
	const prizeCard = 10
	player := NewPlayer("mozart", strategy.NEXT_CARD)

	// test
	player.WinsRound(prizeCard)

	playerStats := player.PlayerStats
	ok := (playerStats.GameTotal == prizeCard) && (playerStats.NumRoundsWon == 1)

	if !ok {
		t.Errorf("WinsRound error")
	}
}

func TestMakeOffer(t *testing.T) {
	cases := []struct {
		inCards       []int
		prizeCard     int
		maxCard       int
		expectedOffer int
		expectedCards []int
	}{
		{[]int{4, 3, 2}, 10, 12, 4, []int{3, 2}},
	}

	for _, c := range cases {
		player := NewPlayer("mozart", strategy.NEXT_CARD)
		hand := NewHand(c.inCards)
		player.SetHand(hand)

		// test
		player.MakeOffer(c.prizeCard, c.maxCard)
		result := player.GetOffer()
		resultCards := player.GetCardsForTesting()

		ok := (result == c.expectedOffer) && (len(c.expectedCards) == len(resultCards))

		if ok {
			for i := 0; i < len(resultCards); i++ {
				if ok && resultCards[i] != c.expectedCards[i] {
					ok = false
				}
			}
		}

		if !ok {
			t.Errorf("MakeOffer offer: %v, expectedOffer: %v, cards: %v expectedCards: %v",
				result, c.expectedOffer, resultCards, c.expectedCards)
		}
	}
}

func TestSolicitOffers(t *testing.T) {
	const prizeCard = 10
	const maxCard = 12

	p1 := BuildPlayerForTesting("beethoven", strategy.NEXT_CARD, []int{3, 2, 1})
	p2 := BuildPlayerForTesting("chopin", strategy.NEXT_CARD, []int{4, 5, 6})
	p3 := BuildPlayerForTesting("mozart", strategy.NEXT_CARD, []int{9, 8, 7})

	players := []Player{p1, p2, p3}

	// test
	SolicitOffers(prizeCard, players, maxCard)

	// TODO: figure out how to use test-assert library (see README.md)
	ok := (players[0].offer == 3) && (players[1].offer == 4) && (players[2].offer == 9)

	if !ok {
		t.Errorf("SolicitOffers() error")
	}
}
