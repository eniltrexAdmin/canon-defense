package game

type HardcodedLevelGenerator struct{}

func (lg HardcodedLevelGenerator) Generate(level int) CanonTDGame {

	bg := NewBattleGround(5, 9, 5)
	cd := NewCanonDeck(bg)
	mt := NewMonsterTeam(bg)

	mt.addMonster(8, 3, SkeletonTemplate())
	mt.addMonster(4, 3, SkeletonTemplate())

	return CanonTDGame{
		Battleground: bg,
		CanonDeck:    cd,
		MonsterTeam:  mt,
	}
}
