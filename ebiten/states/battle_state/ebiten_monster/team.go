package ebiten_monster

import (
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_background"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_canon"
	"canon-tower-defense/game"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type EbitenMonsterTeam struct {
	visibleMonsters []*EbitenMonster
	advanceStep     float64
	lateralStep     float64
}

func NewEbitenMonsterTeam(g *game.CanonTDGame) EbitenMonsterTeam {
	f := func(event EbitenMonsterHitEvent) {
		g.HitMonster(event.Canon, event.Monster)
	}

	availableWidth := constants.ScreenWidth / int(g.MonsterTeam.Battleground.Columns)
	availableHeight := int(ebiten_background.BattleGroundHeight) / int(g.MonsterTeam.Battleground.VisibleRows)
	var visibleMonsters []*EbitenMonster

	centerXSpace := availableWidth / 2
	tileCenterPointX := float64(centerXSpace)

	centerYSpace := availableHeight / 2
	tileSCenterPointY := float64(centerYSpace)

	for _, mig := range g.MonsterTeam.Monsters {
		println(fmt.Printf("adding monster in %d\n", mig.CurrentVisibleRow))

		posX := float64(availableWidth*int(mig.CurrentColumn)) + tileCenterPointX
		posY := float64(availableHeight*int(mig.CurrentVisibleRow)) + tileSCenterPointY

		monster := NewEbitenMonster(mig, posX, posY, f)
		visibleMonsters = append(visibleMonsters, monster)
	}

	return EbitenMonsterTeam{
		visibleMonsters: visibleMonsters,
		advanceStep:     float64(availableHeight),
		lateralStep:     float64(availableWidth),
	}
}

func (emt *EbitenMonsterTeam) Draw(screen *ebiten.Image) {
	for _, visibleMonsters := range emt.visibleMonsters {
		visibleMonsters.Draw(screen)
	}
}

func (emt *EbitenMonsterTeam) Update() {
	for _, visibleMonsters := range emt.visibleMonsters {
		visibleMonsters.Update()
	}
}

func (emt *EbitenMonsterTeam) DeckFiring(bullets []*ebiten_canon.EbitenCanonBullet) {
	for _, visibleMonsters := range emt.visibleMonsters {
		visibleMonsters.DeckFiring(bullets)
	}
}

func (emt *EbitenMonsterTeam) Advance() {
	for _, visibleMonsters := range emt.visibleMonsters {
		visibleMonsters.Attack(emt.advanceStep, emt.lateralStep)
	}
}

func (emt *EbitenMonsterTeam) AreAttacking() bool {
	for _, visibleMonster := range emt.visibleMonsters {
		if visibleMonster.IsAttacking() {
			return true
		}
	}
	return false
}

func (emt *EbitenMonsterTeam) AreAlive() bool {
	for _, visibleMonster := range emt.visibleMonsters {
		if visibleMonster.monster.IsAlive() {
			return true
		}
	}
	return false
}

func (emt *EbitenMonsterTeam) ReachedGameOver() bool {
	for _, visibleMonster := range emt.visibleMonsters {
		if visibleMonster.ReachedGameOver() {
			return true
		}
	}
	return false
}
