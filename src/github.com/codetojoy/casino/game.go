
package casino

import (
    "github.com/codetojoy/config"
    "github.com/codetojoy/player"
)

func PlayGame(config config.Config, players []player.Player) {

    table := Deal(config, players)

    for i := 0; i < config.NumCardsPerHand; i++ {
        PlayRound(config, table)
    }
}
