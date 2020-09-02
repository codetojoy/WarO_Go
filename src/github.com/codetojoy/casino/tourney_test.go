
package casino

import (
    "testing"

    "github.com/codetojoy/player"
)

func TestDetermineTourneyWinner(t *testing.T) {
    p1 := player.NewPlayer("beethoven")
    p2 := player.NewPlayer("chopin")
    p3 := player.NewPlayer("mozart")

    p1.PlayerStats.NumGamesWon = 0
    p2.PlayerStats.NumGamesWon = 3
    p3.PlayerStats.NumGamesWon = 4

    players := []player.Player{p1, p2, p3}

    // test
    result := determineTourneyWinner(players)

    // TODO: figure out how to use test-assert library (see README.md)
    ok := result.GetName() == p3.GetName()

    if ! ok {
        t.Errorf("determineTourneyWinner() error actual: %v expected %v", result.GetName(), p3.GetName())
    }
}
