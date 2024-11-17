package game

import "fmt"

type battlegroundMovement uint8

type MonsterTemplate struct {
	name           string
	healthPoints   canonDamage
	rowMovement    battlegroundMovement
	columnMovement battlegroundMovement
}

func SkeletonTemplate() MonsterTemplate {
	return MonsterTemplate{
		name:           "Skeleton",
		healthPoints:   1,
		columnMovement: 0,
		rowMovement:    1,
	}
}

type Monster struct {
	Monster           MonsterTemplate
	name              string
	maxLife           canonDamage
	healthPoints      canonDamage
	rowMovement       battlegroundMovement
	columnMovement    battlegroundMovement
	CurrentColumn     BattleGroundColumn
	CurrentRow        BattleGroundRow
	CurrentVisibleRow BattleGroundRow
	hitHistory        map[Turn]MonsterHit
}

func (m *Monster) alive() bool {
	return m.healthPoints > 0
}

func (m *Monster) Hit(c *Canon, turn Turn) {
	if m.CurrentVisibleRow == NoVisibleRow {
		panic(fmt.Sprintf("hitting not visible monster?"))
	}

	if _, exists := m.hitHistory[turn]; exists {
		return
	}
	m.healthPoints -= c.Damage
	m.hitHistory[turn] = MonsterHit{
		Damage: c.Damage,
		Turn:   turn,
	}
}

type MonsterHit struct {
	Damage canonDamage
	Turn   Turn
}
