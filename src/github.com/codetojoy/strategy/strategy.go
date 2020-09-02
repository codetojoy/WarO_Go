
package strategy

type Strategy struct {
}

func (strategy *Strategy) SelectCard(prizeCard int, cards []int, maxCard int) int {
    return cards[0]
}
