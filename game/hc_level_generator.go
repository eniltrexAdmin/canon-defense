package game

type LevelGenerator struct{}

func (lg LevelGenerator) Generate() battleground {

	var monsterFormation [][]*Monster

	for i := 1; i <= 10; i++ {
		monsterRow := make([]*Monster, 5)
		monsterFormation = append(monsterFormation, monsterRow)
	}

	skeleton := Skeleton()
	monsterFormation[10][3] = &skeleton

	return battleground{
		columns:  defaultBattleGroundSize,
		rows:     BattleGroundLength(len(monsterFormation)),
		monsters: monsterFormation,
	}
}
