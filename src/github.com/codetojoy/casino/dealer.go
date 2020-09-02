
package casino

import (
    "math/rand"
    "time"

    "github.com/codetojoy/config"
    "github.com/codetojoy/player"
)

func Deal(config config.Config, players []player.Player) Table {
    deck := newDeck(config.NumCards)
    deck = shuffle(deck)
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

func newDeck(numCards int) []int {
    deck := []int{}

    for i := 1; i <= numCards; i++ {
        deck = append(deck, i)
    }

    return deck
}

func shuffle(cards []int) []int {
    rand.Seed(time.Now().UnixNano())

    rand.Shuffle(len(cards), func(i, j int) {
        cards[i], cards[j] = cards[j], cards[i]
    })

    return cards
}

func partition(cards []int, numCardsPerHand int) []player.Hand {
    result := []player.Hand{}
    thisHand := player.NewHandNoCards()

    for i, card := range cards {
        thisHand.DealCard(card)

        if (i > 0) && ((i + 1) % numCardsPerHand == 0) {
            result = append(result, thisHand)
            thisHand = player.NewHandNoCards()
        }
    }

    return result
}
