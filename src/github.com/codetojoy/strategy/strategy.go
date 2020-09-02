
package strategy

import (
    "strings"
)

type Strategy interface {
    SelectCard(prizeCard int, cards []int, maxCard int) int
}

func BuildStrategy(which string) Strategy {
    var result Strategy

    const NEAREST_CARD = "NearestCard"
    const NEXT_CARD = "NextCard"

    if strings.EqualFold(NEXT_CARD, which) {
        result = nextCard{}
    } else if strings.EqualFold(NEAREST_CARD, which) {
        result = nearestCard{}
    }

    return result
}

type nextCard struct {
}

func (nc nextCard) SelectCard(prizeCard int, cards []int, maxCard int) int {
    return cards[0]
}

type nearestCard struct {
}

func absValue(x int) int {
    result := x

    if x < 0 {
        result = -1 * x
    }

    return result
}

func (nc nearestCard) SelectCard(prizeCard int, cards []int, maxCard int) int {
    result := 0
    bestAbsValue := maxCard * 10

    for _, card := range cards {
        thisAbsValue := absValue(card - prizeCard)

        if thisAbsValue < bestAbsValue {
            bestAbsValue = thisAbsValue
            result = card
        }
    }

    return result
}
