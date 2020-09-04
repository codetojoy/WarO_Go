
package casino

import (
    "github.com/codetojoy/config"
    "github.com/codetojoy/player"
)

// 'proper' vs 'testing'
type TestDealer struct {
    deckProvider DeckProvider
}

func NewTestDealer(deckProvider DeckProvider) TestDealer {
    return TestDealer{deckProvider: deckProvider}
}

func (testDealer TestDealer) deal(config config.Config, players []player.Player) Table {
    deck := testDealer.deckProvider.newDeck(config.NumCards)
    // we don't want random functionality in tests:
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
    table := newTable(kitty, players)

    return table
}
