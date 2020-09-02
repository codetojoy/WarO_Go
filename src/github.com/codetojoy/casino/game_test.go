
package casino

import (
    "testing"

    "github.com/codetojoy/player"
)

func TestDetermineGameWinner(t *testing.T) {
    p1 := player.NewPlayer("beethoven")
    p2 := player.NewPlayer("chopin")
    p3 := player.NewPlayer("mozart")

    p1.PlayerStats.GameTotal = 18
    p2.PlayerStats.GameTotal = 30
    p3.PlayerStats.GameTotal = 0

    players := []player.Player{p1, p2, p3}

    // test
    result := determineGameWinner(players)

    // TODO: figure out how to use test-assert library (see README.md)
    ok := result.GetName() == p2.GetName()

    if ! ok {
        t.Errorf("determineGameWinner() error actual: %v expected %v", result.GetName(), p2.GetName())
    }
}
