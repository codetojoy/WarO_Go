package casino

type TestDeckProvider struct {
	deck []int
}

func (deckProvider TestDeckProvider) newDeck(numCards int) []int {
	return deckProvider.deck
}
