
package casino

import (
    "fmt"

    "github.com/codetojoy/config"
    "github.com/codetojoy/player"
)

func PlayTourney(config config.Config, players []player.Player, dealer Dealer) {
    for i := 0; i < config.NumGames; i++ {
        playGame(config, players, dealer)
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
    metric := player.ByNumGamesWon
    return player.FindHighestByMetric(players, metric)
}
