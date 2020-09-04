
package strategy

import (
    "testing"
)

func TestNextCard(t *testing.T) {
    const prizeCard = 10
    cards := []int{4,6,8}
    const maxCard = 12
    strategy := BuildStrategy(NEXT_CARD)

    // test
    result := strategy.SelectCard(prizeCard, cards, maxCard)

    // TODO: figure out how to use test-assert library (see README.md)
    expected := 4
    ok := (result == expected)

    if ! ok {
        t.Errorf("NextCard actual: %d expected: %d", result, expected)
    }
}

func TestNearestCard(t *testing.T) {
    const prizeCard = 10
    cards := []int{4,6,9,12}
    const maxCard = 12
    strategy := BuildStrategy(NEAREST_CARD)

    // test
    result := strategy.SelectCard(prizeCard, cards, maxCard)

    // TODO: figure out how to use test-assert library (see README.md)
    expected := 9
    ok := (result == expected)

    if ! ok {
        t.Errorf("NearestCard actual: %d expected: %d", result, expected)
    }
}

func TestMaxCard(t *testing.T) {
    const prizeCard = 10
    cards := []int{4,11,9}
    const maxCard = 12
    strategy := BuildStrategy(MAX_CARD)

    // test
    result := strategy.SelectCard(prizeCard, cards, maxCard)

    // TODO: figure out how to use test-assert library (see README.md)
    expected := 11
    ok := (result == expected)

    if ! ok {
        t.Errorf("MaxCard actual: %d expected: %d", result, expected)
    }
}

func TestMinCard(t *testing.T) {
    const prizeCard = 10
    cards := []int{4,6,9,12}
    const maxCard = 12
    strategy := BuildStrategy(MIN_CARD)

    // test
    result := strategy.SelectCard(prizeCard, cards, maxCard)

    // TODO: figure out how to use test-assert library (see README.md)
    expected := 4
    ok := (result == expected)

    if ! ok {
        t.Errorf("MinCard actual: %d expected: %d", result, expected)
    }
}

func TestHybrid_Max(t *testing.T) {
    const prizeCard = 10
    cards := []int{4,6,9,12}
    const maxCard = 12
    strategy := BuildStrategy(HYBRID)

    // test
    result := strategy.SelectCard(prizeCard, cards, maxCard)

    // TODO: figure out how to use test-assert library (see README.md)
    expected := 12
    ok := (result == expected)

    if ! ok {
        t.Errorf("Hybrid actual: %d expected: %d", result, expected)
    }
}

func TestHybrid_Min(t *testing.T) {
    const prizeCard = 2
    cards := []int{4,6,9,12}
    const maxCard = 12
    strategy := BuildStrategy(HYBRID)

    // test
    result := strategy.SelectCard(prizeCard, cards, maxCard)

    // TODO: figure out how to use test-assert library (see README.md)
    expected := 4
    ok := (result == expected)

    if ! ok {
        t.Errorf("Hybrid actual: %d expected: %d", result, expected)
    }
}

func TestIsLegalPick(t *testing.T) {
    cases := []struct {
        inCards []int
        pick int
        expected bool
    }{
        {[]int{1,2,3}, 2, true},
        {[]int{1,2,3}, 44, false},
    }

    consoleCard := consoleCard{}

    for _, c := range cases {
        // test
        result := consoleCard.isLegalPick(c.inCards, c.pick)

        ok := (result == c.expected)

        if ! ok {
            t.Errorf("isLegalPick() == %v, expected %v", result, c.expected)
        }
    }
}
