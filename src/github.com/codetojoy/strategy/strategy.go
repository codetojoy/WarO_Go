
package strategy

import (
    "strings"
)

type Strategy interface {
    SelectCard(prizeCard int, cards []int, max int) int
}

const HYBRID = "Hybrid"
const NEAREST_CARD = "Nearest_Card"
const NEXT_CARD = "Next_Card"
const MAX_CARD = "Max_Card"
const MIN_CARD = "Min_Card"

func BuildStrategy(which string) Strategy {
    var result Strategy

    if strings.EqualFold(HYBRID, which) {
        result = hybridCard{}
    } else if strings.EqualFold(NEXT_CARD, which) {
        result = nextCard{}
    } else if strings.EqualFold(NEAREST_CARD, which) {
        result = nearestCard{}
    } else if strings.EqualFold(MAX_CARD, which) {
        result = maxCard{}
    } else if strings.EqualFold(MIN_CARD, which) {
        result = minCard{}
    }

    return result
}

type nextCard struct {
}

func (nc nextCard) SelectCard(prizeCard int, cards []int, max int) int {
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

func (nc nearestCard) SelectCard(prizeCard int, cards []int, max int) int {
    result := 0
    bestAbsValue := max * 10

    for _, card := range cards {
        thisAbsValue := absValue(card - prizeCard)

        if thisAbsValue < bestAbsValue {
            bestAbsValue = thisAbsValue
            result = card
        }
    }

    return result
}

type maxCard struct {
}

func (mc maxCard) SelectCard(prizeCard int, cards []int, maxCard int) int {
    result := 0

    for _, card := range cards {
        if card > result {
            result = card
        }
    }

    return result
}

type minCard struct {
}

func (mc minCard) SelectCard(prizeCard int, cards []int, maxCard int) int {
    result := maxCard * 10

    for _, card := range cards {
        if card < result {
            result = card
        }
    }

    return result
}

type hybridCard struct {
}

func (hc hybridCard) SelectCard(prizeCard int, cards []int, max int) int {

    var s Strategy

    if prizeCard > (max / 2) {
        s = maxCard{}
    } else {
        s = minCard{}
    }

    result := s.SelectCard(prizeCard, cards, max)

    return result
}
