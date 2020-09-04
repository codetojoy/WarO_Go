package casino

import (
	"fmt"

	"github.com/codetojoy/config"
	"github.com/codetojoy/player"
)

func playRound(config config.Config, table *Table) {
	playRoundSimple(&table.kitty, table.players, config.MaxCard)
	fmt.Printf("---\n")
}

// "simple" in the sense of params with simpler levels of nesting
func playRoundSimple(kitty *player.Hand, players []player.Player, maxCard int) {
	prizeCard := kitty.TakeCard()
	player.SolicitOffers(prizeCard, players, maxCard)

	for index := range players {
		player := &players[index]
		fmt.Printf("TRACER playRound %v bids %v on %v\n", player.GetName(), player.GetOffer(), prizeCard)
	}

	winner := determineRoundWinner(players)
	winner.WinsRound(prizeCard)
}

func determineRoundWinner(players []player.Player) *player.Player {
	metric := player.ByOffer
	return player.FindHighestByMetric(players, metric)
}
