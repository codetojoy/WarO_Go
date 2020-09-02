
package strategy

import (
    "strings"
)

type Strategy interface {
    SelectCard(prizeCard int, cards []int, maxCard int) int
}

func BuildStrategy(which string) Strategy {
    result := nextCard{}

    const NEXT_CARD = "NextCard"

    if strings.EqualFold(NEXT_CARD, which) {
       // no-op
    }

    return result
}

type nextCard struct {
}

func (nc nextCard) SelectCard(prizeCard int, cards []int, maxCard int) int {
    return cards[0]
}
