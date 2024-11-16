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

	for i := game.BattleGroundRow(0); i < monsterTeam.Battleground.VisibleRows; i++ {
		for j := game.BattleGroundColumn(0); j < monsterTeam.Battleground.Columns; j++ {
			centerXSpace := availableWidth / 2
			tileCenterPointX := float64(centerXSpace)

			centerYSpace := availableHeight / 2
			tileSCenterPointY := float64(centerYSpace)

			posX := float64(availableWidth*int(j)) + tileCenterPointX
			posY := float64(availableHeight*int(i)) + tileSCenterPointY

			if monsterTeam.MonstersInField[j][i] != nil {
				monster := NewEbitenMonster(monsterTeam.MonstersInField[j][i], posX, posY)
				visibleMonsters = append(visibleMonsters, &monster)
			}
		}
	}

	return ebitenMonsterTeam{
		visibleMonsters:  visibleMonsters,
		monsterAttacking: false,
		advanceStep:      float64(availableHeight),
	}
}

func (ecd *ebitenMonsterTeam) draw(screen *ebiten.Image) {
	for _, visibleMonsters := range ecd.visibleMonsters {
		visibleMonsters.draw(screen)
	}
}

func (ecd *ebitenMonsterTeam) update(bullets []*ebitenCanonBullet) {
	for _, visibleMonsters := range ecd.visibleMonsters {
		visibleMonsters.update(bullets)
	}
}

func (ecd *ebitenMonsterTeam) monsterAdvancePositions(numPositions int) {
	ecd.monsterAttacking = true
	for _, visibleMonsters := range ecd.visibleMonsters {
		visibleMonsters.setDestination(ecd.advanceStep * float64(numPositions))
	}
}

func (ecd *ebitenMonsterTeam) updateAttack() {
	monsterMoving := false
	for _, visibleMonsters := range ecd.visibleMonsters {
		visibleMonsters.updateAttack()
		if visibleMonsters.isMoving {
			monsterMoving = true
		}
	}
	if !monsterMoving {
		ecd.monsterAttacking = false
	}
}
