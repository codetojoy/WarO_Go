
package strategy

import (
    "testing"
)

func TestNextCard(t *testing.T) {
    const prizeCard = 10
    cards := []int{4,6,8}
    const maxCard = 12
    strategy := BuildStrategy("nextcard")

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
    strategy := BuildStrategy("nearestcard")

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
    strategy := BuildStrategy("maxcard")

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
    strategy := BuildStrategy("mincard")

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
    strategy := BuildStrategy("hybrid")

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
    strategy := BuildStrategy("hybrid")

    // test
    result := strategy.SelectCard(prizeCard, cards, maxCard)

    // TODO: figure out how to use test-assert library (see README.md)
    expected := 4
    ok := (result == expected)

    if ! ok {
        t.Errorf("Hybrid actual: %d expected: %d", result, expected)
    }
}
