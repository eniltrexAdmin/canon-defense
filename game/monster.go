package game

import "fmt"

type battlegroundMovement uint8

type MonsterTemplate struct {
	name           string
	healthPoints   canonDamage
	rowMovement    battlegroundMovement
	columnMovement battlegroundMovement
}

func SkeletonTemplate(life int) MonsterTemplate {
	return MonsterTemplate{
		name:           "Skeleton",
		healthPoints:   canonDamage(life),
		columnMovement: 0,
		rowMovement:    1,
	}
}

type Monster struct {
	name              string
	MaxLife           canonDamage
	HealthPoints      canonDamage
	RowMovement       battlegroundMovement
	columnMovement    battlegroundMovement
	CurrentColumn     BattleGroundColumn
	CurrentRow        BattleGroundRow
	CurrentVisibleRow BattleGroundRow
	hitHistory        map[Turn]MonsterHit
}

func newMonsterInBattleGround(
	bg Battleground,
	column BattleGroundColumn,
	row BattleGroundRow,
	m MonsterTemplate,
) Monster {
	bg.checkIndexPosition(row, column)
	return Monster{
		name:              m.name,
		MaxLife:           m.healthPoints,
		HealthPoints:      m.healthPoints,
		RowMovement:       m.rowMovement,
		columnMovement:    m.columnMovement,
		CurrentColumn:     column,
		CurrentRow:        row,
		CurrentVisibleRow: bg.toVisibleRow(row),
		hitHistory:        make(map[Turn]MonsterHit),
	}
}

func (m *Monster) IsAlive() bool {
	return m.HealthPoints > 0
}

func (m *Monster) Hit(c *Canon, turn Turn) {
	if m.CurrentVisibleRow == NoVisibleRow {
		panic(fmt.Sprintf("hitting not visible monster?"))
	}

	if _, exists := m.hitHistory[turn]; exists {
		return
	}
	m.HealthPoints -= c.Damage
	m.hitHistory[turn] = MonsterHit{
		Damage: c.Damage,
		Turn:   turn,
	}
}

type MonsterHit struct {
	Damage canonDamage
	Turn   Turn
}
