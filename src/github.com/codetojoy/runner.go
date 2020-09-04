
package main

import (
    "fmt"
    "os"

    "github.com/codetojoy/casino"
    "github.com/codetojoy/config"
    "github.com/codetojoy/player"
)

func main() {
    numArgs := len(os.Args)

    if numArgs > 1 {
        configFile := os.Args[1]

        config := config.NewConfigFromFile(configFile)
        fmt.Printf("TRACER config: %v\n", config.String())

        players := player.BuildPlayers(config)

        deckProvider := casino.SimpleDeckProvider{}
        dealer := casino.NewProperDealer(deckProvider)
        
        casino.PlayTourney(config, players, dealer)
    } else {
        fmt.Println("Usage.")
    }

    fmt.Println("Ready.")
}
