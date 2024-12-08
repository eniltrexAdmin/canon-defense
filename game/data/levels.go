package data

import "canon-tower-defense/game"

func Level1() game.CanonTDGame {
	bg := game.NewBattleGround(5, 9, 5)
	cd := game.NewCanonDeck(bg)
	mt := game.NewMonsterTeam(bg)

	mt.AddMonster(1, 2, SlimeTemplate())
	mt.AddMonster(2, 0, SlimeTemplate())
	mt.AddMonster(4, 4, SlimeTemplate())
	mt.AddMonster(3, 4, SlimeTemplate())
	mt.AddMonster(7, 2, SlimeTemplate())
	mt.AddMonster(8, 3, BeholderTemplate())
	mt.AddMonster(8, 0, BeholderTemplate())
	mt.AddMonster(8, 2, SlimeTemplate())

	return game.NewGame(bg, cd, mt)
}

func Level2() game.CanonTDGame {
	bg := game.NewBattleGround(5, 9, 5)
	cd := game.NewCanonDeck(bg)
	mt := game.NewMonsterTeam(bg)

	mt.AddMonster(8, 3, BeholderTemplate())
	mt.AddMonster(2, 4, BeholderTemplate())
	mt.AddMonster(4, 3, BeholderTemplate())
	mt.AddMonster(3, 0, BeholderTemplate())
	mt.AddMonster(1, 2, BeholderTemplate())

	return game.NewGame(bg, cd, mt)
}

func Level3() game.CanonTDGame {
	bg := game.NewBattleGround(5, 9, 5)
	cd := game.NewCanonDeck(bg)
	mt := game.NewMonsterTeam(bg)

	mt.AddMonster(8, 3, LizardFolkTemplate())
	mt.AddMonster(5, 4, PurpleWormTemplate())
	mt.AddMonster(4, 3, LizardFolkTemplate())
	mt.AddMonster(2, 0, PurpleWormTemplate())
	mt.AddMonster(2, 2, PurpleWormTemplate())

	return game.NewGame(bg, cd, mt)
}

func Level4() game.CanonTDGame {
	bg := game.NewBattleGround(5, 9, 5)
	cd := game.NewCanonDeck(bg)
	mt := game.NewMonsterTeam(bg)

	mt.AddMonster(4, 2, DragonTemplate())
	mt.AddMonster(4, 0, LizardFolkTemplate())
	mt.AddMonster(4, 1, BeholderTemplate())
	mt.AddMonster(4, 3, PurpleWormTemplate())

	mt.AddMonster(3, 1, MedusaTemplate())
	mt.AddMonster(3, 2, SlimeTemplate())
	mt.AddMonster(3, 3, SerpentFlyTemplate())

	return game.NewGame(bg, cd, mt)
}
