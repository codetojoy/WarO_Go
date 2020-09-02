
package casino

import (
    "github.com/codetojoy/config"
    "github.com/codetojoy/player"
)

func PlayTourney(config config.Config, players []player.Player) {
    for i := 0; i < config.NumGames; i++ {
        PlayGame(config, players)
    }
}
