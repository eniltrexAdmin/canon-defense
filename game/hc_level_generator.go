package game

type HardcodedLevelGenerator struct{}

func (lg HardcodedLevelGenerator) Generate(level int) CanonTDGame {

	bg := NewBattleGround(5, 9, 5)
	cd := NewCanonDeck(bg)
	mt := NewMonsterTeam(bg)

	mt.AddMonster(8, 3, Skeleton())
	mt.AddMonster(4, 3, Skeleton())

	return CanonTDGame{
		Battleground: bg,
		CanonDeck:    cd,
		MonsterTeam:  mt,
	}
}
