
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
    result := GetBids(prizeCard, players, maxCard)

    // TODO: figure out how to use test-assert library (see README.md)
    ok := len(result) == len(players)
    ok = ok && (result[0].Offer == 3)
    ok = ok && (result[1].Offer == 4)
    ok = ok && (result[2].Offer == 9) 

    if ! ok {
        t.Errorf("GetBids() error")
    }
}
