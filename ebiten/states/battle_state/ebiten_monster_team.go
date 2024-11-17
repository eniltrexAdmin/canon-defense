package battle_state

import (
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type ebitenMonsterTeam struct {
	game             *game.CanonTDGame
	visibleMonsters  []*ebitenMonster
	monsterAttacking bool
	advanceStep      float64
}

func NewEbitenMonsterTeam(g game.CanonTDGame) ebitenMonsterTeam {
	availableWidth := constants.ScreenWidth / int(g.MonsterTeam.Battleground.Columns)
	availableHeight := int(BattleGroundHeight) / int(g.MonsterTeam.Battleground.VisibleRows)
	var visibleMonsters []*ebitenMonster

	centerXSpace := availableWidth / 2
	tileCenterPointX := float64(centerXSpace)

	centerYSpace := availableHeight / 2
	tileSCenterPointY := float64(centerYSpace)

	for _, mig := range g.MonsterTeam.Monsters {
		// at some point we might have get rid of visible rows all together.
		if mig.CurrentVisibleRow != game.NoVisibleRow {
			posX := float64(availableWidth*int(mig.CurrentColumn)) + tileCenterPointX
			posY := float64(availableHeight*int(mig.CurrentVisibleRow)) + tileSCenterPointY

			monster := NewEbitenMonster(mig, posX, posY)
			visibleMonsters = append(visibleMonsters, &monster)
		}
	}

	return ebitenMonsterTeam{
		visibleMonsters:  visibleMonsters,
		monsterAttacking: false,
		advanceStep:      float64(availableHeight),
		game:             &g,
	}
}

func (emt *ebitenMonsterTeam) draw(screen *ebiten.Image) {
	for _, visibleMonsters := range emt.visibleMonsters {
		visibleMonsters.draw(screen)
	}
}

// BATTLE STATE

func (emt *ebitenMonsterTeam) update() {
	for _, visibleMonsters := range emt.visibleMonsters {
		visibleMonsters.update()
	}
}

// DECK FIRING STATE

func (emt *ebitenMonsterTeam) updateDeckFiring(bullets []*ebitenCanonBullet) {
	for _, visibleMonsters := range emt.visibleMonsters {
		visibleMonsters.updateDeckFiring(bullets)
		if visibleMonsters.isHit {
			emt.game.HitMonster(
				&visibleMonsters.hittingBullet.canon.canon,
				visibleMonsters.monster,
			)
		}
	}
}

// ATTACK STATE

func (emt *ebitenMonsterTeam) monsterAdvancePositions(numPositions int) {
	emt.monsterAttacking = true
	for _, visibleMonsters := range emt.visibleMonsters {
		visibleMonsters.setDestination(emt.advanceStep * float64(numPositions))
	}
}

func (emt *ebitenMonsterTeam) updateAttack() {
	monsterMoving := false
	for _, visibleMonsters := range emt.visibleMonsters {
		visibleMonsters.updateAttack()
		if visibleMonsters.isMoving {
			monsterMoving = true
		}
	}
	if !monsterMoving {
		emt.monsterAttacking = false
	}
}
