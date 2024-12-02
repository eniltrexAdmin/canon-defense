package constants

import "canon-tower-defense/game/player"

const (
	ScreenWidth  = 450
	ScreenHeight = 800
)

type GameContext struct {
	Player player.Player
}

var GlobalContext = &GameContext{}
