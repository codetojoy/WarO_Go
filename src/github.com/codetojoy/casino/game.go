
package casino

import (
    "fmt"

    "github.com/codetojoy/config"
    "github.com/codetojoy/player"
)

func PlayGame(config config.Config, players []player.Player) {
    fmt.Printf("TRACER play game\n")

    table := Deal(config, players)

    for i := 1; i <= config.NumCardsPerHand; i++ {
        PlayRound(config, table)
    }
}
