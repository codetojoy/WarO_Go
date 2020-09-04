
package casino

import (
    "fmt"
    "strings"

    "github.com/codetojoy/player"
)

type Table struct {
    kitty player.Hand
    players []player.Player
}

func newTable(kitty player.Kitty, players []player.Player) Table {
    return Table{kitty: kitty, players: players}
}

func (table *Table) String() string {
    result := strings.Builder{}

    result.WriteString(fmt.Sprintf("kitty: %v\n", table.kitty.String()))
    for _, player := range table.players {
        result.WriteString(fmt.Sprintf("%v\n", player.String()))
    }

    return result.String()
}
