
package player

import (
    "testing"
)

func TestDealCard(t *testing.T) {
    cases := []struct {
        inCards []int
        expected []int
    }{
        {[]int{1,2,3}, []int{1,2,3}},
    }

    for _, c := range cases {
        hand := NewHandNoCards()

        // test
        for _, card := range c.inCards {
            hand.DealCard(card)
        }
        result := hand.GetCards()

        ok := len(c.expected) == len(result)

        if ok {
            for i := 0; i < len(result); i++ {
                if ok && result[i] != c.expected[i] {
                    ok = false
                }
            }
        }

        if ! ok {
            t.Errorf("DealCard() == %v, expected %v", result, c.expected)
        }
    }
}

func TestRemoveCard(t *testing.T) {
    cases := []struct {
        inCards []int
        target int
        expected []int
    }{
        {[]int{1,2,3}, 2, []int{1,3}},
    }

    for _, c := range cases {
        hand := NewHand(c.inCards)

        // test
        hand.RemoveCard(c.target)
        result := hand.GetCards()

        ok := len(c.expected) == len(result)

        if ok {
            for i := 0; i < len(result); i++ {
                if ok && result[i] != c.expected[i] {
                    ok = false
                }
            }
        }

        if ! ok {
            t.Errorf("RemoveCard() == %v, expected %v", result, c.expected)
        }
    }
}

func TestTakeCard(t *testing.T) {
    cases := []struct {
        inCards []int
        expected []int
    }{
        {[]int{1,2,3}, []int{2,3}},
    }

    for _, c := range cases {
        hand := NewHand(c.inCards)

        // test
        hand.TakeCard()
        result := hand.GetCards()

        ok := len(c.expected) == len(result)

        if ok {
            for i := 0; i < len(result); i++ {
                if ok && result[i] != c.expected[i] {
                    ok = false
                }
            }
        }

        if ! ok {
            t.Errorf("TakeCard() == %v, expected %v", result, c.expected)
        }
    }
}
