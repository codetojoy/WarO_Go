
package casino

import (
    "github.com/codetojoy/config"
    "github.com/codetojoy/player"
)

// 'proper' vs 'testing'
type TestDealer struct {
}

func (testDealer TestDealer) deal(config config.Config, players []player.Player) Table {
    deck := newDeck(config.NumCards)
    // deck = shuffle(deck)
    hands := partition(deck, config.NumCardsPerHand)
    var kitty player.Kitty

    for i, hand := range hands {
        if i > 0 {
            players[i-1].SetHand(hand)
        } else {
            kitty = hand
        }
    }
    table := NewTable(kitty, players)

    return table
}
