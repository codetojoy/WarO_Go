
package casino

import (
    "fmt"
    
    "github.com/codetojoy/config"
    "github.com/codetojoy/player"
)

func PlayTourney(config config.Config, players []player.Player) {
    fmt.Printf("TRACER play tourney\n")
    for i := 0; i < config.NumGames; i++ {
        PlayGame(config, players)
    }
}
