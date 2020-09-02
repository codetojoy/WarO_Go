
package casino

import (
    "testing"

    "github.com/codetojoy/player"
)

func TestGetBids(t *testing.T) {
    const prizeCard = 10
    const maxCard = 12

    p1 := player.BuildPlayerForTesting("beethoven", []int{3,2,1})
    p2 := player.BuildPlayerForTesting("chopin", []int{4,5,6})
    p3 := player.BuildPlayerForTesting("mozart", []int{9,8,7})

    players := []player.Player{p1, p2, p3}

    // test
    result := getBids(prizeCard, players, maxCard)

    // TODO: figure out how to use test-assert library (see README.md)
    ok := len(result) == len(players)
    ok = ok && (result[0].Offer == 3) && (result[0].Player.GetName() == "beethoven")
    ok = ok && (result[1].Offer == 4) && (result[1].Player.GetName() == "chopin")
    ok = ok && (result[2].Offer == 9) && (result[2].Player.GetName() == "mozart")

    if ! ok {
        t.Errorf("getBids() error")
    }
}

func TestPlayRound(t *testing.T) {
    const maxCard = 12

    kitty := player.NewHand([]int{10,11,12})
    p1 := player.BuildPlayerForTesting("beethoven", []int{3,2,1})
    p2 := player.BuildPlayerForTesting("chopin", []int{4,5,6})
    p3 := player.BuildPlayerForTesting("mozart", []int{9,8,7})

    // TODO: this makes a copy of Player
    players := []player.Player{p1, p2, p3}

    // test
    playRound(&kitty, players, maxCard)

    // TODO: figure out how to use test-assert library (see README.md)
    ok := len(kitty.GetCards()) == 2
    ok = ok && len(players[0].GetCardsForTesting()) == 2
    ok = ok && len(players[1].GetCardsForTesting()) == 2
    ok = ok && len(players[2].GetCardsForTesting()) == 2

    if ! ok {
        t.Errorf("playRound() error")
    }
}
