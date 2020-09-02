
package casino

import (
    "testing"
)

func TestNewDeck(t *testing.T) {
    cases := []struct {
        in int
        expected []int
    }{
        {7, []int{1,2,3,4,5,6,7}},
    }

    for _, c := range cases {
        // test
        result := newDeck(c.in)

        ok := len(c.expected) == len(result)

        if ok {
            for i := 0; i < len(result); i++ {
                if ok && result[i] != c.expected[i] {
                    ok = false
                }
            }
        }

        if ! ok {
            t.Errorf("newDeck() == %v, expected %v", result, c.expected)
        }
    }
}

func TestShuffle(t *testing.T) {
    cases := []struct {
        in []int
        sumCheck int
    }{
        {[]int{1,2,3,4,5,6,7}, 28},
    }

    for _, c := range cases {
        // test
        result := shuffle(c.in)

        total := 0
        ok := len(c.in) == len(result)

        if ok {
            for _, card := range result {
                total += card
            }

            ok = (total == c.sumCheck)
        }

        if ! ok {
            t.Errorf("shuffle() == %v, sumCheck %v", result, c.sumCheck)
        }
    }
}

func TestPartition(t *testing.T) {
    cases := []struct {
        cards []int
        numCardsPerHand int
        sumCheck int
    }{
        {[]int{1,2,3,4,5,6,7,8,9}, 3, 45},
    }

    for _, c := range cases {
        // test
        result := partition(c.cards, c.numCardsPerHand)

        sum := 0
        count := 0
        ok := true

        for _, hand := range result {
            cards := hand.GetCards()
            for _, card := range cards {
                sum += card
                count++
            }
            if len(cards) != c.numCardsPerHand {
                ok = false
            }
        }

        if ! (ok && (sum == c.sumCheck) && (count == len(c.cards))) {
            t.Errorf("partition() == %v, sumCheck %v", result, c.sumCheck)
        }
    }
}
