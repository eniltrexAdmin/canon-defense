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
	//monsterAttacking bool
	//advanceStep      float64
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
		//monsterAttacking: false,
		//advanceStep:      float64(availableHeight),
		game: &g,
	}
}

func (emt *EbitenMonsterTeam) Draw(screen *ebiten.Image) {
	for _, visibleMonsters := range emt.visibleMonsters {
		visibleMonsters.Draw(screen)
	}
}

// BATTLE STATE

func (emt *EbitenMonsterTeam) Update() {
	for _, visibleMonsters := range emt.visibleMonsters {
		visibleMonsters.Update()
	}
}

// DECK FIRING STATE

func (emt *EbitenMonsterTeam) UpdateDeckFiring(bullets []*ebiten_canon.EbitenCanonBullet) {
	for _, visibleMonsters := range emt.visibleMonsters {
		visibleMonsters.UpdateDeckFiring(bullets)
		//if visibleMonsters.isHit {
		//	emt.game.HitMonster(
		//		&visibleMonsters.hittingBullet.canon.canon,
		//		visibleMonsters.monster,
		//	)
		//	visibleMonsters.LifeLine.SetCurrentLife(int(visibleMonsters.monster.HealthPoints))
		//}
	}
}

// ATTACK STATE

//func (emt *ebitenMonsterTeam) monsterAdvancePositions(numPositions int) {
//	emt.monsterAttacking = true
//	for _, visibleMonsters := range emt.visibleMonsters {
//		visibleMonsters.setDestination(emt.advanceStep * float64(numPositions))
//	}
//}
//

func (emt *EbitenMonsterTeam) UpdateAttack() {
	//monsterMoving := false
	//for _, visibleMonsters := range emt.visibleMonsters {
	//	visibleMonsters.updateAttack()
	//	if visibleMonsters.isMoving {
	//		monsterMoving = true
	//	}
	//}
	//if !monsterMoving {
	//	emt.monsterAttacking = false
	//}
}
