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
		return Level10()
	case 9:
		return Level9()
	case 8:
		return Level8()
	case 7:
		return Level7()
	case 6:
		return Level6()
	case 5:
		return Level5()
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
