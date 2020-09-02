
package casino

import (
    "fmt"

    "github.com/codetojoy/config"
    "github.com/codetojoy/player"
)

func PlayRound(config config.Config, table *Table) {
    playRound(&table.kitty, table.players, config.MaxCard)
    fmt.Printf("---\n")
}

func playRound(kitty *player.Hand, players []player.Player, maxCard int) {
    prizeCard := kitty.GetCards()[0]
    // fmt.Printf("TRACER playRound kitty: %v prizeCard: %d\n", kitty.String(), prizeCard)
    bids := getBids(prizeCard, players, maxCard)
    for _, bid := range bids {
        fmt.Printf("TRACER playRound %v bids %v on %v\n", bid.Player.GetName(), bid.Offer, bid.PrizeCard)
    }
    kitty.RemoveCard(prizeCard)
}

func getBids(prizeCard int, players []player.Player, maxCard int) []player.Bid {
    bids := []player.Bid{}

    // `index, player = range` will give a 'value' copy, so be careful!
    for index := range players {
        thisPlayer := &players[index]
        offer := thisPlayer.GetOffer(prizeCard, maxCard)
        bid := player.Bid{Offer: offer, PrizeCard: prizeCard, Player: thisPlayer}
        bids = append(bids, bid)
    }

    return bids
}
