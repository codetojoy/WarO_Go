package casino

type DeckProvider interface {
	newDeck(numCards int) []int
}

type SimpleDeckProvider struct {
}

func (deckProvider SimpleDeckProvider) newDeck(numCards int) []int {
	deck := []int{}

	for i := 1; i <= numCards; i++ {
		deck = append(deck, i)
	}

	return deck
}
