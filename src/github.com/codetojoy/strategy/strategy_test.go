
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
