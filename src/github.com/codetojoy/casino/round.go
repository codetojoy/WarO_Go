
package casino

import (
    "fmt"

    "github.com/codetojoy/config"
    "github.com/codetojoy/player"
)

func PlayRound(config config.Config, table Table) {
    for _, prizeCard := range table.kitty.GetCards() {
        bids := GetBids(prizeCard, table.players, config.MaxCard)
        for _, bid := range bids {
            fmt.Printf("TRACER play round %v bids %v on %v\n", bid.Player.GetName(), bid.Offer, bid.PrizeCard)
        }
        table.kitty.RemoveCard(prizeCard)
        fmt.Printf("---\n")
    }
}

func GetBids(prizeCard int, players[] player.Player, maxCard int) []player.Bid {
    bids := []player.Bid{}

    for _, thisPlayer := range players {
        offer := thisPlayer.GetOffer(prizeCard, maxCard)
        bid := player.Bid{Offer: offer, PrizeCard: prizeCard, Player: thisPlayer}
        bids = append(bids, bid)
    }

    return bids
}
