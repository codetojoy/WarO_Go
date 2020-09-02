
package player

import (
    "testing"
)

func TestWinsRound(t *testing.T) {
    const prizeCard = 10
    player := NewPlayer("mozart")

    // test
    player.WinsRound(prizeCard)

    playerStats := player.PlayerStats
    ok := (playerStats.GameTotal == prizeCard) && (playerStats.NumRoundsWon == 1)

    if ! ok {
        t.Errorf("Wins Round error")
    }
}

func TestGetOffer(t *testing.T) {
    cases := []struct {
        inCards []int
        prizeCard int
        maxCard int
        expectedOffer int
        expectedCards []int
    }{
        {[]int{4,3,2}, 10, 12, 4, []int{3,2}},
    }

    for _, c := range cases {
        player := NewPlayer("mozart")
        hand := NewHand(c.inCards)
        player.SetHand(hand)

        // test
        result := player.GetOffer(c.prizeCard, c.maxCard)
        resultCards := player.GetCardsForTesting()

        ok := (result == c.expectedOffer) && (len(c.expectedCards) == len(resultCards))

        if ok {
            for i := 0; i < len(resultCards); i++ {
                if ok && resultCards[i] != c.expectedCards[i] {
                    ok = false
                }
            }
        }

        if ! ok {
            t.Errorf("GetOffer offer: %v, expectedOffer: %v, cards: %v expectedCards: %v",
                        result, c.expectedOffer, resultCards, c.expectedCards)
        }
    }
}
