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
	game            *game.CanonTDGame
	visibleMonsters []*EbitenMonster
	advanceStep     float64
}

func NewEbitenMonsterTeam(g game.CanonTDGame) EbitenMonsterTeam {
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
		if mig.CurrentVisibleRow != game.NoVisibleRow {
			posX := float64(availableWidth*int(mig.CurrentColumn)) + tileCenterPointX
			posY := float64(availableHeight*int(mig.CurrentVisibleRow)) + tileSCenterPointY

			monster := NewEbitenMonster(mig, posX, posY, &g)
			visibleMonsters = append(visibleMonsters, monster)
		}
	}

	return EbitenMonsterTeam{
		visibleMonsters: visibleMonsters,
		game:            &g,
		advanceStep:     float64(availableHeight),
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
		visibleMonsters.Advance(emt.advanceStep)
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
