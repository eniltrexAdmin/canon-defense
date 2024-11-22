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
	VisibleMonsters []*EbitenMonster
	advanceStep     float64
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
		// at some point we might have get rid of visible rows all together.
		//if mig.CurrentVisibleRow != game.NoVisibleRow {
		posX := float64(availableWidth*int(mig.CurrentColumn)) + tileCenterPointX
		posY := float64(availableHeight*int(mig.CurrentVisibleRow)) + tileSCenterPointY

		monster := NewEbitenMonster(mig, posX, posY, f)
		visibleMonsters = append(visibleMonsters, monster)
		//}
	}

	return EbitenMonsterTeam{
		VisibleMonsters: visibleMonsters,
		advanceStep:     float64(availableHeight),
	}
}

func (emt *EbitenMonsterTeam) Draw(screen *ebiten.Image) {
	for _, visibleMonsters := range emt.VisibleMonsters {
		visibleMonsters.Draw(screen)
	}
}

func (emt *EbitenMonsterTeam) Update() {
	for _, visibleMonsters := range emt.VisibleMonsters {
		visibleMonsters.Update()
	}
}

func (emt *EbitenMonsterTeam) DeckFiring(bullets []*ebiten_canon.EbitenCanonBullet) {
	for _, visibleMonsters := range emt.VisibleMonsters {
		visibleMonsters.DeckFiring(bullets)
	}
}

func (emt *EbitenMonsterTeam) Advance() {
	for _, visibleMonsters := range emt.VisibleMonsters {
		visibleMonsters.Attack(emt.advanceStep)
	}
}

func (emt *EbitenMonsterTeam) AreAttacking() bool {
	for _, visibleMonster := range emt.VisibleMonsters {
		if visibleMonster.IsAttacking() {
			return true
		}
	}
	return false
}
