
package casino

import (
    "fmt"

    "github.com/codetojoy/config"
    "github.com/codetojoy/player"
)

func PlayTourney(config config.Config, players []player.Player) {
    for i := 0; i < config.NumGames; i++ {
        PlayGame(config, players)
    }

    winner := determineTourneyWinner(players)

    fmt.Println("**************")
    fmt.Printf("TRACER %v WINS tourney!\n", winner.GetName())
    for index := range players {
        player := &players[index]
        player.NewGame() // clear last game stats
        fmt.Printf("%v\n", player.String())
    }
}

func determineTourneyWinner(players []player.Player) *player.Player {
    result := &players[0]
    bestTotal := result.PlayerStats.NumGamesWon

    for index := range players {
        player := &players[index]
        thisTotal := player.PlayerStats.NumGamesWon
        if thisTotal > bestTotal {
            bestTotal = thisTotal
            result = player
        }
    }

    return result
}
