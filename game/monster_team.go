package game

import "fmt"

//9 Rows 5  columns
//  0 1 2 3 4
//9
//8
//7            6 Visible rows
//6                              visibleRow := total Visible mt.LevelRows - gameRow - 1
//5           0                     x = 6 - 5 - 1  = 0
//4           1                     x = 6 - 4 - 1 = 1
//3           2                     x = 6 - 3 -1 = 2
//2           3
//1           4
//0           5                     x = 6 -0 -1 = 5

type MonsterTeam struct {
	Monsters        map[BattleGroundColumn]map[BattleGroundRow]*Monster
	MonstersInField map[BattleGroundColumn]map[BattleGroundRow]*Monster
	Battleground    Battleground
}

func NewMonsterTeam(bg Battleground) MonsterTeam {
	m := make(map[BattleGroundColumn]map[BattleGroundRow]*Monster, bg.Columns)

	for c := BattleGroundColumn(0); c < bg.Columns; c++ {
		m[c] = make(map[BattleGroundRow]*Monster, bg.Rows)
	}

	monstersInField := make(map[BattleGroundColumn]map[BattleGroundRow]*Monster, bg.Columns)
	for c := BattleGroundColumn(0); c < bg.Columns; c++ {
		monstersInField[c] = make(map[BattleGroundRow]*Monster, bg.VisibleRows)
	}

	return MonsterTeam{
		Monsters:        m,
		MonstersInField: monstersInField,
		Battleground:    bg,
	}
}

func (mt *MonsterTeam) AddMonster(indexRow, indexColumn int, m Monster) {
	mt.Battleground.checkIndexPosition(indexRow, indexColumn)
	bgColumn := BattleGroundColumn(indexColumn)
	bgRow := BattleGroundRow(indexRow)
	mt.Monsters[bgColumn][bgRow] = &m

	visibleRow, err := mt.toVisibleRow(bgRow)
	if err == nil {
		mt.MonstersInField[bgColumn][visibleRow] = &m
	}
}

func (mt *MonsterTeam) toRealRow(visibleRow BattleGroundRow) BattleGroundRow {
	return mt.Battleground.Rows - (visibleRow + 1)
}

func (mt *MonsterTeam) toVisibleRow(gameRow BattleGroundRow) (BattleGroundRow, error) {
	if gameRow < mt.Battleground.Rows-mt.Battleground.VisibleRows ||
		gameRow >= mt.Battleground.VisibleRows {
		return 0, fmt.Errorf("gameRow %d is out of the visible range", gameRow)
	}

	visibleRow := mt.Battleground.VisibleRows - gameRow - 1
	return visibleRow, nil
}
