
package casino

import (
    "testing"

    "github.com/codetojoy/player"
    "github.com/codetojoy/strategy"
)


func TestPlayRoundSimple(t *testing.T) {
    const maxCard = 12

    kitty := player.NewHand([]int{10,11,12})
    p1 := player.BuildPlayerForTesting("beethoven", strategy.NEXT_CARD, []int{3,2,1})
    p2 := player.BuildPlayerForTesting("chopin", strategy.NEXT_CARD, []int{4,5,6})
    p3 := player.BuildPlayerForTesting("mozart", strategy.NEXT_CARD, []int{9,8,7})

    // TODO: this makes a copy of Player
    players := []player.Player{p1, p2, p3}

    // test
    playRoundSimple(&kitty, players, maxCard)

    // TODO: figure out how to use test-assert library (see README.md)
    ok := len(kitty.GetCards()) == 2
    ok = ok && len(players[0].GetCardsForTesting()) == 2
    ok = ok && len(players[1].GetCardsForTesting()) == 2
    ok = ok && len(players[2].GetCardsForTesting()) == 2

    if ! ok {
        t.Errorf("playRound() error")
    }
}

func TestDetermineRoundWinner(t *testing.T) {
    const prizeCard = 10
    const maxCard = 12

    p1 := player.BuildPlayerForTesting("beethoven", strategy.NEXT_CARD, []int{3,2,1})
    p2 := player.BuildPlayerForTesting("chopin", strategy.NEXT_CARD, []int{4,5,6})
    p3 := player.BuildPlayerForTesting("mozart", strategy.NEXT_CARD, []int{9,8,7})
    players := []player.Player{p1, p2, p3}

    player.SolicitOffers(prizeCard, players, maxCard)

    // test
    result := determineRoundWinner(players)

    // TODO: figure out how to use test-assert library (see README.md)
    ok :=  (result.GetName() == "mozart") && (result.GetOffer() == 9)

    if ! ok {
        t.Errorf("determineRoundWinner() error")
    }
}
