package game

import (
	"fmt"
)

type BattlegroundMovement uint8

type MonsterTemplate struct {
	Name           string
	HealthPoints   CanonDamage
	RowMovement    BattlegroundMovement
	ColumnMovement BattlegroundMovement
}

type Monster struct {
	Name              string
	MaxLife           CanonDamage
	HealthPoints      CanonDamage
	RowMovement       BattlegroundMovement
	columnMovement    BattlegroundMovement
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
		Name:              m.Name,
		MaxLife:           m.HealthPoints,
		HealthPoints:      m.HealthPoints,
		RowMovement:       m.RowMovement,
		columnMovement:    m.ColumnMovement,
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
		println(fmt.Sprintf("hitting not visible monster? name: %s, column: %d, row %d", m.Name, m.CurrentVisibleRow, m.CurrentRow))
		return
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
	Damage CanonDamage
	Turn   Turn
}
