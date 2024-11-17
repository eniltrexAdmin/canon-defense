package game

// at some point just having one index should suffice:
// satting the "goal" like 6 or 7, and total rows
// so a monster i from -5 to goal, when it's on 0 or more it's visible

//9 Rows 5  columns
//  0 1 2 3 4
//9
//8
//7            6 Visible rows
//6                              VisibleRow := total Visible mt.LevelRows - gameRow - 1
//5           0                     x = 6 - 5 - 1  = 0
//4           1                     x = 6 - 4 - 1 = 1
//3           2                     x = 6 - 3 -1 = 2
//2           3
//1           4
//0           5                     x = 6 -0 -1 = 5

type MonsterTeam struct {
	Battleground Battleground
	Monsters     []*Monster
}

func NewMonsterTeam(bg Battleground) MonsterTeam {
	return MonsterTeam{
		Battleground: bg,
	}
}

func (mt *MonsterTeam) addMonster(indexRow, indexColumn int, m MonsterTemplate) {
	bgColumn := BattleGroundColumn(indexColumn)
	bgRow := BattleGroundRow(indexRow)
	monsterInBg := newMonsterInBattleGround(mt.Battleground, bgColumn, bgRow, m)
	mt.Monsters = append(mt.Monsters, &monsterInBg)
}

//func (mt *MonsterTeam) monsterInColumn(c BattleGroundColumn) []*Monster {
//	m := make([]*Monster, 0)
//	for _, monsterInBg := range mt.Monsters {
//		if monsterInBg.CurrentVisibleRow != NoVisibleRow && monsterInBg.CurrentColumn == c {
//			m = append(m, monsterInBg)
//		}
//	}
//	return m
//}

//func (mt *MonsterTeam) DamageMonsters(c *Canon, canonPosition BattleGroundColumn) {
//	for _, monster := range mt.monsterInColumn(canonPosition) {
//		monster.hit(c.Damage)
//	}
//}

//func (mt *MonsterTeam) advance() {
//	for _, monsterInBg := range mt.MonstersInBattleground {
//		monsterInBg.Row = monsterInBg.Row - 1
//		monsterInBg.VisibleRow = mt.Battleground.toVisibleRow(monsterInBg.Row)
//	}
//}
