package player

import (
	"fmt"
)

type PlayerStats struct {
	GameTotal    int
	NumRoundsWon int
	NumGamesWon  int
}

func (ps *PlayerStats) String() string {
	return fmt.Sprintf("[%d, %d, %d]", ps.GameTotal, ps.NumRoundsWon, ps.NumGamesWon)
}
