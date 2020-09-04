
package json

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

type JsonPlayer struct {
    Name string `json:"name"`
    Strategy string `json:"strategy"`
}

type JsonConfig struct {
    NumCards int `json:"num_cards"`
    NumGames int `json:"num_games"`
    IsVerbose bool `json:"is_verbose"`
    Players []JsonPlayer `json:"players"`
}

func (jsonConfig *JsonConfig) Log() {
    fmt.Printf("numCards: %d\n", jsonConfig.NumCards)
    fmt.Printf("numGames: %d\n", jsonConfig.NumGames)
    fmt.Printf("isVerbose: %v\n", jsonConfig.IsVerbose)

    for i, _ := range jsonConfig.Players {
        player := &jsonConfig.Players[i]
        fmt.Printf("player name: %v strategy: %v\n", player.Name, player.Strategy)
    }
}

func BuildJsonConfig(jsonFileName string) *JsonConfig {
    // from https://tutorialedge.net/golang/parsing-json-with-golang/
    jsonFile, err := os.Open(jsonFileName)
    defer jsonFile.Close()

    if err != nil {
        log.Fatal(err)
    }

    data, _ := ioutil.ReadAll(jsonFile)

    var jsonConfig JsonConfig

    json.Unmarshal(data, &jsonConfig)

    return &jsonConfig
}
