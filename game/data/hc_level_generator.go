package data

import "canon-tower-defense/game"

type HardcodedLevelGenerator struct{}

func (lg HardcodedLevelGenerator) Generate(level int) game.CanonTDGame {

	switch level {
	case 4:
		return Level4()
	case 3:
		return Level3()
	case 2:
		return Level2()
	default:
		return Level4()
	}
}
