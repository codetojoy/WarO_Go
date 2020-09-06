package strategy

import (
	"fmt"
	"strings"
	"testing"
)

func TestBuildUrl(t *testing.T) {
	const prizeCard = 10
	const maxCard = 12
	const port = 6160
	const mode = "min"
	const context = "/waro/strategy"
	cards := []int{3, 7, 11}

	urlStr := fmt.Sprintf("http://localhost:%d%v", port, context)

	// test
	result := buildUrl(port, urlStr, prizeCard, cards, maxCard, mode)

	// TODO: figure out how to use test-assert library (see README.md)
	ok := strings.Contains(result, "http://localhost:6160/waro/strategy?")
	ok = ok && strings.Contains(result, "cards=3&cards=7&cards=11")
	ok = ok && strings.Contains(result, "max_card=12")
	ok = ok && strings.Contains(result, "prize_card=10")
	ok = ok && strings.Contains(result, "max_card=12")
	ok = ok && strings.Contains(result, "mode=min")

	if !ok {
		t.Errorf("parseJson\n actual: %v \n", result)
	}
}

func TestParseJson(t *testing.T) {
	jsonStr := `{"card": 5150, "message": ""}`

	// test
	result := parseJson(jsonStr)

	// TODO: figure out how to use test-assert library (see README.md)
	expected := 5150
	ok := (result == expected)

	if !ok {
		t.Errorf("parseJson actual: %d expected: %d", result, expected)
	}
}
