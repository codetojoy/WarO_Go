
package casino

import (
    "fmt"

    "github.com/codetojoy/config"
    "github.com/codetojoy/player"
)

func playGame(config config.Config, players []player.Player) {

    table := Deal(config, players)

    for index := range players {
        player := &players[index]
        player.NewGame()
    }

    for i := 0; i < config.NumCardsPerHand; i++ {
        fmt.Printf("\nTRACER PlayGame round %d table: \n%v\n\n", i, table.String())
        playRound(config, &table)
    }

    winner := determineGameWinner(players)
    winner.WinsGame()

    fmt.Printf("\nTRACER PlayGame %v WINS table: \n%v\n\n", winner.GetName(), table.String())
}

func determineGameWinner(players []player.Player) *player.Player {
    result := &players[0]
    bestTotal := result.PlayerStats.GameTotal

    for index := range players {
        player := &players[index]
        thisTotal := player.PlayerStats.GameTotal
        if thisTotal > bestTotal {
            bestTotal = thisTotal
            result = player
        }
    }

    return result
}
