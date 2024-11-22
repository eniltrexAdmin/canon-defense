package game

type HardcodedLevelGenerator struct{}

func (lg HardcodedLevelGenerator) Generate(level int) CanonTDGame {

	bg := NewBattleGround(5, 9, 5)
	cd := NewCanonDeck(bg)
	mt := NewMonsterTeam(bg)

	mt.addMonster(8, 3, SkeletonTemplate(15))
	mt.addMonster(2, 4, SkeletonTemplate(1))
	mt.addMonster(4, 3, SkeletonTemplate(2))
	mt.addMonster(3, 0, SkeletonTemplate(1))
	mt.addMonster(1, 2, SkeletonTemplate(1))

	return CanonTDGame{
		Battleground: bg,
		CanonDeck:    cd,
		MonsterTeam:  mt,
		playerTurn:   true,
	}
}
