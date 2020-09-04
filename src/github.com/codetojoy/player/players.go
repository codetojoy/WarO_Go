
package player

import (
    "github.com/codetojoy/config"
)

func BuildPlayers(config config.Config) []Player {
    result := []Player{}

    for _, configPlayer := range config.Players {
        name := configPlayer.Name
        whichStrategy := configPlayer.WhichStrategy
        player := NewPlayer(name, whichStrategy)
        result = append(result, player)
    }

    return result
}

func SolicitOffers(prizeCard int, players []Player, maxCard int) {
    // `index, player = range` will give a 'value' copy, so be careful!
    for index := range players {
        thisPlayer := &players[index]
        thisPlayer.MakeOffer(prizeCard, maxCard)
    }
}

func FindHighestByMetric(players []Player, f Metric) *Player {
    result := &players[0]
    bestValue := f(result)

    for index := range players {
        player := &players[index]
        thisValue := f(player)

        if thisValue > bestValue {
            bestValue = thisValue
            result = player
        }
    }

    return result
}
