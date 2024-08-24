package game

type HardcodedLevelGenerator struct{}

func (lg HardcodedLevelGenerator) Generate(level int) Battleground {

	var monsterFormation [][]*Monster

	for i := 1; i <= 10; i++ {
		monsterRow := make([]*Monster, 5)
		monsterFormation = append(monsterFormation, monsterRow)
	}

	skeleton := Skeleton()
	monsterFormation[9][3] = &skeleton

	return Battleground{
		Columns:  defaultBattleGroundSize,
		Rows:     BattleGroundLength(len(monsterFormation)),
		monsters: monsterFormation,
	}
}
