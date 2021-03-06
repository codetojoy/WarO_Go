// Package strategy contains various strategies (see Strategy in Design Patterns).
package strategy

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Strategy interface {
	SelectCard(ch chan int, prizeCard int, cards []int, max int)
}

const API_REMOTE = "api_remote"
const CONSOLE = "console"
const HYBRID = "hybrid"
const NEAREST_CARD = "nearest_card"
const NEXT_CARD = "next_card"
const MAX_CARD = "max_card"
const MIN_CARD = "min_card"

func BuildStrategy(which string, url string) Strategy {
	var result Strategy

	if strings.EqualFold(API_REMOTE, which) {
		result = apiRemoteCard{url: url}
	} else if strings.EqualFold(CONSOLE, which) {
		result = consoleCard{}
	} else if strings.EqualFold(HYBRID, which) {
		result = hybridCard{}
	} else if strings.EqualFold(NEXT_CARD, which) {
		result = nextCard{}
	} else if strings.EqualFold(NEAREST_CARD, which) {
		result = nearestCard{}
	} else if strings.EqualFold(MAX_CARD, which) {
		result = maxCard{}
	} else if strings.EqualFold(MIN_CARD, which) {
		result = minCard{}
	}

	return result
}

type nextCard struct {
}

func (nc nextCard) SelectCard(ch chan int, prizeCard int, cards []int, _ int) {
	ch <- cards[0]
}

type nearestCard struct {
}

func absValue(x int) int {
	result := x

	if x < 0 {
		result = -1 * x
	}

	return result
}

func (nc nearestCard) SelectCard(ch chan int, prizeCard int, cards []int, max int) {
	result := 0
	bestAbsValue := max * 10

	for _, card := range cards {
		thisAbsValue := absValue(card - prizeCard)

		if thisAbsValue < bestAbsValue {
			bestAbsValue = thisAbsValue
			result = card
		}
	}

	ch <- result
}

type maxCard struct {
}

func (mc maxCard) SelectCard(ch chan int, prizeCard int, cards []int, _ int) {
	result := 0

	for _, card := range cards {
		if card > result {
			result = card
		}
	}

	ch <- result
}

type minCard struct {
}

func (mc minCard) SelectCard(ch chan int, prizeCard int, cards []int, max int) {
	result := max * 10

	for _, card := range cards {
		if card < result {
			result = card
		}
	}

	ch <- result
}

type hybridCard struct {
}

func (hc hybridCard) SelectCard(ch chan int, prizeCard int, cards []int, max int) {
	var s Strategy

	if prizeCard > (max / 2) {
		s = maxCard{}
	} else {
		s = minCard{}
	}

	s.SelectCard(ch, prizeCard, cards, max)
}

type consoleCard struct {
}

func (cc consoleCard) SelectCard(ch chan int, prizeCard int, cards []int, _ int) {
	ok := false

	var card int
	var err error

	for !ok {
		card, err = cc.getInput(prizeCard, cards)

		ok = (err == nil) && cc.isLegalPick(cards, card)

		if !ok {
			fmt.Printf("illegal pick!\n")
		}
	}

	ch <- card
}

func (cc consoleCard) isLegalPick(cards []int, pick int) bool {
	result := false

	for _, card := range cards {
		if (!result) && (card == pick) {
			result = true
		}
	}

	return result
}

func (cc consoleCard) getInput(prizeCard int, cards []int) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("------------\n")
	fmt.Printf("prize card: %d\nyour hand: %v\n\n", prizeCard, cards)
	fmt.Print("your choice: ")

	var card int
	var err error
	s, e1 := reader.ReadString('\n')

	if e1 == nil {
		card, err = strconv.Atoi(strings.TrimSpace(s))
	} else {
		err = e1
	}

	return card, err
}
