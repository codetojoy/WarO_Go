
package main

import (
    "fmt"

    "github.com/codetojoy/casino"
    "github.com/codetojoy/config"
    "github.com/codetojoy/player"
)

func main() {
    config := config.NewConfig()
    fmt.Printf("TRACER config: %v\n", config)
    players := player.BuildPlayers(config)

    casino.PlayTourney(config, players)

    fmt.Println("Ready.")
    /*
    numArgs := len(os.Args)

    if numArgs > 1 {
        mode := os.Args[1]

        if mode == "process50" {
            n := 50
            processN(n)
        } else if mode == "search" {
            x := 121
            search(x)
        } else {
            fmt.Println("ERROR: unknown mode")
            os.Exit(1)
        }
    } else {
        fmt.Println("USAGE: requires mode param")
    }
    1G*/
}
