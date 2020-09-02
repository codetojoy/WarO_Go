
package casino

import (
    "fmt"

    "github.com/codetojoy/config"
    "github.com/codetojoy/player"
)

func PlayGame(config config.Config, players []player.Player) {

    table := Deal(config, players)

    for i := 0; i < config.NumCardsPerHand; i++ {
        fmt.Printf("\nTRACER PlayGame round %d table: \n%v\n\n", i, table.String())
        PlayRound(config, &table)
    }
}
