// Package player handles player state/functionality.
package player

import (
    "fmt"
    "strings"
)

type Hand struct {
    cards []int
}

type Kitty = Hand

func NewHandNoCards() Hand {
    return NewHand([]int{})
}

func NewHand(cards []int) Hand {
    return Hand{cards: cards}
}

func (hand *Hand) DealCard(card int) {
    hand.cards = append(hand.cards, card)
}

// TODO: this is pretty crazy and not idiomatic
func (hand *Hand) RemoveCard(card int) {
    cards := []int{}

    for _, thisCard := range hand.cards {
        if thisCard != card {
            cards = append(cards, thisCard)
        }
    }

    hand.cards = cards
}

func (hand *Hand) GetCards() []int {
    return hand.cards
}

func (hand *Hand) String() string {
    result := strings.Builder{}

    numCards := len(hand.cards)

    result.WriteString("[")
    for i, card := range hand.cards {
        result.WriteString(fmt.Sprintf("%d", card))

        isLastCard := (i == (numCards - 1))
        if ! isLastCard {
            result.WriteString(" ")
        }
    }
    result.WriteString("]")

    return result.String()
}
