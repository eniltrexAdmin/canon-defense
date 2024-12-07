package data

import "canon-tower-defense/game"

func Level1() game.CanonTDGame {
	bg := game.NewBattleGround(5, 9, 5)
	cd := game.NewCanonDeck(bg)
	mt := game.NewMonsterTeam(bg)

	mt.AddMonster(1, 2, Beholder())

	return game.NewGame(bg, cd, mt)
}

func Level2() game.CanonTDGame {
	bg := game.NewBattleGround(5, 9, 5)
	cd := game.NewCanonDeck(bg)
	mt := game.NewMonsterTeam(bg)

	mt.AddMonster(8, 3, Beholder())
	mt.AddMonster(2, 4, Beholder())
	mt.AddMonster(4, 3, Beholder())
	mt.AddMonster(3, 0, Beholder())
	mt.AddMonster(1, 2, Beholder())

	return game.NewGame(bg, cd, mt)
}

func Level3() game.CanonTDGame {
	bg := game.NewBattleGround(5, 9, 5)
	cd := game.NewCanonDeck(bg)
	mt := game.NewMonsterTeam(bg)

	mt.AddMonster(8, 3, LizardFolkTemplate())
	mt.AddMonster(5, 4, PurpleWorm())
	mt.AddMonster(4, 3, LizardFolkTemplate())
	mt.AddMonster(2, 0, PurpleWorm())
	mt.AddMonster(2, 2, PurpleWorm())

	return game.NewGame(bg, cd, mt)
}

func Level4() game.CanonTDGame {
	bg := game.NewBattleGround(5, 9, 5)
	cd := game.NewCanonDeck(bg)
	mt := game.NewMonsterTeam(bg)

	mt.AddMonster(4, 2, DragonTemplate())

	return game.NewGame(bg, cd, mt)
}
