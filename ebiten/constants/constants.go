package constants

import (
	"canon-tower-defense/game/session"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	ScreenWidth  = 450
	ScreenHeight = 800
)

type GameContext struct {
	Session session.Session
}

var GlobalContext = &GameContext{}

var AudioContext = audio.NewContext(44100)
