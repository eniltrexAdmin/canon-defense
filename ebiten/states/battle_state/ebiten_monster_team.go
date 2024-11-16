package battle_state

import (
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type ebitenMonsterTeam struct {
	visibleMonsters  []*ebitenMonster
	monsterAttacking bool
	advanceStep      float64
}

func NewEbitenMonsterTeam(monsterTeam game.MonsterTeam) ebitenMonsterTeam {
	availableWidth := constants.ScreenWidth / int(monsterTeam.Battleground.Columns)
	availableHeight := int(BattleGroundHeight) / int(monsterTeam.Battleground.VisibleRows)
	var visibleMonsters []*ebitenMonster

	centerXSpace := availableWidth / 2
	tileCenterPointX := float64(centerXSpace)

	centerYSpace := availableHeight / 2
	tileSCenterPointY := float64(centerYSpace)

	for _, mig := range monsterTeam.MonstersInBattleground {
		// at some point we might have get rid of visible rows all together.
		if mig.VisibleRow != game.NoVisibleRow {
			posX := float64(availableWidth*int(mig.Column)) + tileCenterPointX
			posY := float64(availableHeight*int(mig.VisibleRow)) + tileSCenterPointY

			monster := NewEbitenMonster(&mig.Monster, posX, posY)
			visibleMonsters = append(visibleMonsters, &monster)
		}
	}

	return ebitenMonsterTeam{
		visibleMonsters:  visibleMonsters,
		monsterAttacking: false,
		advanceStep:      float64(availableHeight),
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
