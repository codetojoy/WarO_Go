// Package strategy contains various strategies (see Strategy in Design Patterns).
package strategy

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const CARDS_PARAM = "cards"
const MAX_CARD_PARAM = "max_card"
const PRIZE_CARD_PARAM = "prize_card"

const MODE_PARAM = "mode"
const MODE_MAX = "max"
const MODE_MIN = "min"

type apiRemoteCard struct {
    url string
}

type ApiRemoteResult struct {
	Card    int
	Message string
}

func (arc apiRemoteCard) SelectCard(ch chan int, prizeCard int, cards []int, maxCard int) {
	url := buildUrl(arc.url, prizeCard, cards, maxCard)

	jsonStr := callServer(url)
	card := parseJson(jsonStr)

	ch <- card
}

func parseJson(jsonStr string) int {
	var apiRemoteResult ApiRemoteResult
	json.Unmarshal([]byte(jsonStr), &apiRemoteResult)

	return apiRemoteResult.Card
}

func callServer(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println("TRACER response status:", resp.Status)
	result := strings.Builder{}

	// this is weird:
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		result.WriteString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result.String()
}

// func buildUrl(port int, urlStr string, prizeCard int, cards []int, maxCard int, mode string) string {
func buildUrl(urlStr string, prizeCard int, cards []int, maxCard int) string {
	req, err := http.NewRequest("GET", urlStr, nil)

	if err != nil {
		log.Fatal(err)
	}

	query := req.URL.Query()
	// query.Add(MODE_PARAM, mode)
	query.Add(PRIZE_CARD_PARAM, strconv.Itoa(prizeCard))
	query.Add(MAX_CARD_PARAM, strconv.Itoa(maxCard))

	for _, card := range cards {
		query.Add(CARDS_PARAM, strconv.Itoa(card))
	}

	req.URL.RawQuery = query.Encode()

	return req.URL.String()
}
