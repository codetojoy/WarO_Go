
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
    bids := getBids(prizeCard, players, maxCard)
    for _, bid := range bids {
        fmt.Printf("TRACER playRound %v bids %v on %v\n", bid.Player.GetName(), bid.Offer, bid.PrizeCard)
    }
    kitty.RemoveCard(prizeCard)
    winningBid := determineWinner(bids)
    winningBid.Player.WinsRound(prizeCard)
}

func determineWinner(bids []player.Bid) *player.Bid {
    result := &bids[0]
    bestOffer := result.Offer

    // TODO: I don't know yet if Go has a functional 'filter'
    for index := range bids {
        thisBid := &bids[index]
        thisOffer := thisBid.Offer
        if thisOffer > bestOffer {
            bestOffer = thisOffer
            result = thisBid
        }
    }

    return result
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
