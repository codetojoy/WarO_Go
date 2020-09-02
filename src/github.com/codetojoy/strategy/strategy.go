
package strategy

type Strategy interface {
    SelectCard(prizeCard int, cards []int, maxCard int) int
}

type NextCard struct {
}

func (nextCard NextCard) SelectCard(prizeCard int, cards []int, maxCard int) int {
    return cards[0]
}
