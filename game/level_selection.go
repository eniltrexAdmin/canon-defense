package game

import (
	"canon-tower-defense/game/player"
	"fmt"
)

type LevelSelector struct{}

type Levels []bool

const TotalLevels int = 10

func (ls LevelSelector) LevelSelection(pl player.Player) Levels {
	l := make([]bool, TotalLevels)
	for i := 0; i < pl.CurrentLevel; i++ {
		l[i] = true
	}
	fmt.Printf("%t\n", l)
	return l
}
