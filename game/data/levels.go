package data

import "canon-tower-defense/game"

func Level1() (
	game.Battleground,
	game.CanonDeck,
	game.MonsterTeam,
) {
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

	return bg, cd, mt
}

func Level2() (
	game.Battleground,
	game.CanonDeck,
	game.MonsterTeam,
) {
	bg := game.NewBattleGround(5, 9, 5)
	cd := game.NewCanonDeck(bg)
	mt := game.NewMonsterTeam(bg)

	mt.AddMonster(8, 3, BeholderTemplate())
	mt.AddMonster(2, 4, BeholderTemplate())
	mt.AddMonster(4, 3, BeholderTemplate())
	mt.AddMonster(3, 0, BeholderTemplate())
	mt.AddMonster(1, 2, BeholderTemplate())

	return bg, cd, mt
}

func Level3() (
	game.Battleground,
	game.CanonDeck,
	game.MonsterTeam,
) {
	bg := game.NewBattleGround(5, 9, 5)
	cd := game.NewCanonDeck(bg)
	mt := game.NewMonsterTeam(bg)

	mt.AddMonster(8, 3, LizardFolkTemplate())
	mt.AddMonster(5, 4, PurpleWormTemplate())
	mt.AddMonster(4, 3, LizardFolkTemplate())
	mt.AddMonster(6, 0, PurpleWormTemplate())
	mt.AddMonster(6, 2, PurpleWormTemplate())

	return bg, cd, mt
}

func Level4() (
	game.Battleground,
	game.CanonDeck,
	game.MonsterTeam,
) {
	bg := game.NewBattleGround(5, 20, 5)
	cd := game.NewCanonDeck(bg)
	mt := game.NewMonsterTeam(bg)

	mt.AddMonster(1, 2, SlimeTemplate())

	mt.AddMonster(4, 0, LizardFolkTemplate())
	mt.AddMonster(3, 3, SerpentFlyTemplate())
	mt.AddMonster(3, 4, SerpentFlyTemplate())

	mt.AddMonster(10, 0, SerpentFlyTemplate())
	mt.AddMonster(7, 4, SerpentFlyTemplate())
	mt.AddMonster(4, 0, LizardFolkTemplate())
	mt.AddMonster(8, 3, LizardFolkTemplate())

	mt.AddMonster(13, 1, MedusaTemplate())
	mt.AddMonster(14, 3, MedusaTemplate())

	return bg, cd, mt
}

func Level5() (
	game.Battleground,
	game.CanonDeck,
	game.MonsterTeam,
) {
	bg := game.NewBattleGround(5, 20, 5)
	cd := game.NewCanonDeck(bg)
	mt := game.NewMonsterTeam(bg)

	mt.AddMonster(1, 2, SlimeTemplate())
	mt.AddMonster(2, 2, SlimeTemplate())
	mt.AddMonster(3, 2, SlimeTemplate())
	mt.AddMonster(4, 2, SlimeTemplate())
	mt.AddMonster(5, 2, SlimeTemplate())

	mt.AddMonster(3, 1, SlimeTemplate())
	mt.AddMonster(4, 1, SlimeTemplate())
	mt.AddMonster(5, 1, SlimeTemplate())
	mt.AddMonster(6, 1, SlimeTemplate())
	mt.AddMonster(7, 1, SlimeTemplate())

	mt.AddMonster(4, 0, SlimeTemplate())
	mt.AddMonster(5, 0, SlimeTemplate())
	mt.AddMonster(6, 0, SlimeTemplate())
	mt.AddMonster(7, 0, SlimeTemplate())
	mt.AddMonster(8, 0, SlimeTemplate())

	mt.AddMonster(7, 3, MedusaTemplate())

	mt.AddMonster(10, 2, PurpleWormTemplate())
	mt.AddMonster(11, 4, PurpleWormTemplate())

	mt.AddMonster(12, 2, MedusaTemplate())
	mt.AddMonster(12, 3, MedusaTemplate())
	mt.AddMonster(12, 4, MedusaTemplate())

	mt.AddMonster(14, 0, MedusaTemplate())
	mt.AddMonster(14, 1, MedusaTemplate())
	mt.AddMonster(14, 2, MedusaTemplate())

	return bg, cd, mt
}

func Level6() (
	game.Battleground,
	game.CanonDeck,
	game.MonsterTeam,
) {
	bg := game.NewBattleGround(5, 13, 5)
	cd := game.NewCanonDeck(bg)
	mt := game.NewMonsterTeam(bg)

	mt.AddMonster(4, 2, MedusaTemplate())
	mt.AddMonster(5, 3, MedusaTemplate())
	mt.AddMonster(6, 1, MedusaTemplate())
	mt.AddMonster(7, 0, LizardFolkTemplate())
	mt.AddMonster(8, 4, LizardFolkTemplate())

	mt.AddMonster(9, 0, MedusaTemplate())
	mt.AddMonster(9, 1, MedusaTemplate())
	mt.AddMonster(12, 2, MedusaTemplate())
	mt.AddMonster(9, 3, MedusaTemplate())
	mt.AddMonster(12, 4, MedusaTemplate())

	return bg, cd, mt
}

func Level7() (
	game.Battleground,
	game.CanonDeck,
	game.MonsterTeam,
) {
	bg := game.NewBattleGround(5, 20, 5)
	cd := game.NewCanonDeck(bg)
	mt := game.NewMonsterTeam(bg)

	mt.AddMonster(10, 2, DragonTemplate())
	mt.AddMonster(15, 0, DragonTemplate())
	mt.AddMonster(15, 4, DragonTemplate())
	mt.AddMonster(20, 0, DragonTemplate())
	mt.AddMonster(20, 1, DragonTemplate())
	mt.AddMonster(20, 2, DragonTemplate())
	mt.AddMonster(20, 3, DragonTemplate())
	mt.AddMonster(20, 4, DragonTemplate())

	return bg, cd, mt
}

func Level8() (
	game.Battleground,
	game.CanonDeck,
	game.MonsterTeam,
) {
	bg := game.NewBattleGround(5, 31, 5)
	cd := game.NewCanonDeck(bg)
	mt := game.NewMonsterTeam(bg)

	mt.AddMonster(10, 2, DragonTemplate())
	mt.AddMonster(15, 0, DragonTemplate())
	mt.AddMonster(15, 4, DragonTemplate())
	mt.AddMonster(20, 0, DragonTemplate())
	mt.AddMonster(20, 1, DragonTemplate())
	mt.AddMonster(20, 2, DragonTemplate())
	mt.AddMonster(20, 3, DragonTemplate())
	mt.AddMonster(20, 4, DragonTemplate())

	mt.AddMonster(30, 2, SlimeBossTemplate())

	return bg, cd, mt
}

func Level9() (
	game.Battleground,
	game.CanonDeck,
	game.MonsterTeam,
) {
	bg := game.NewBattleGround(5, 31, 10)
	cd := game.NewCanonDeck(bg)
	mt := game.NewMonsterTeam(bg)

	mt.AddMonster(10, 2, AncientDragonTemplate())

	return bg, cd, mt
}

func Level10() (
	game.Battleground,
	game.CanonDeck,
	game.MonsterTeam,
) {
	bg := game.NewBattleGround(6, 21, 10)
	cd := game.NewCanonDeck(bg)
	mt := game.NewMonsterTeam(bg)

	mt.AddMonster(7, 2, AncientDragonTemplate())

	return bg, cd, mt
}
