
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

    deckProvider := SimpleDeckProvider{}

    for _, c := range cases {
        // test
        result := deckProvider.newDeck(c.in)

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
