package data

import "canon-tower-defense/game"

type HardcodedLevelGenerator struct{}

func (lg HardcodedLevelGenerator) Generate(level int) (
	game.Battleground,
	game.CanonDeck,
	game.MonsterTeam,
) {
	switch level {
	case 10:
		return Level4()
	case 9:
		return Level4()
	case 8:
		return Level4()
	case 7:
		return Level4()
	case 6:
		return Level4()
	case 5:
		return Level4()
	case 4:
		return Level4()
	case 3:
		return Level3()
	case 2:
		return Level2()
	default:
		return Level1()
	}
}
