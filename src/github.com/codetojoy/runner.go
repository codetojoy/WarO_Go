
package main

import (
    "fmt"

    "github.com/codetojoy/casino"
    "github.com/codetojoy/config"
    "github.com/codetojoy/player"
)

func main() {
    config := config.NewConfig()
    fmt.Printf("TRACER config: %v\n", config.String())

    players := player.BuildPlayers(config)
    dealer := casino.ProperDealer{}

    casino.PlayTourney(config, players, dealer)

    fmt.Println("Ready.")
}
