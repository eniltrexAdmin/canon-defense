package constants

import "canon-tower-defense/game/session"

const (
	ScreenWidth  = 450
	ScreenHeight = 800
)

type GameContext struct {
	Session session.Session
}

var GlobalContext = &GameContext{}
