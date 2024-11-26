package ebiten_canon

import (
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/game"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

const canonYPlacement float64 = 550

type EbitenCanonDeck struct {
	ebitenCanons        map[int]*ebitenCanon
	deployAreas         map[int]*deployArea
	actionButton        ebitenActionButton
	draggedSprite       *ebiten_sprite.EbitenDraggableSprite
	deployCannonTrigger func(on int)
	moveCannonTrigger   func(from, to int)
}

func NewEbitenCanonDeck(
	g *game.CanonTDGame,
	deployCannonTrigger func(on int),
	moveCannonTrigger func(from, to int),
) EbitenCanonDeck {

	availableWidth := float64(constants.ScreenWidth / g.CanonDeck.CanonCapacity())
	das := make(map[int]*deployArea, g.CanonDeck.CanonCapacity())

	color := ebiten_sprite.RandomColor()
	for formationPlacement := 0; formationPlacement < g.CanonDeck.CanonCapacity(); formationPlacement++ {
		centerX := getCanonCenterX(formationPlacement, g.CanonDeck.CanonCapacity())

		da := NewDeployAreaFromCentralPoint(
			centerX, canonYPlacement,
			availableWidth-20, 60,
			color,
		)
		das[formationPlacement] = &da
	}

	ab := newEbitenActionButton(LoadedImages[RegularCanon], LoadedImages[RegularCanonPedestal], constants.ScreenWidth)

	return EbitenCanonDeck{
		ebitenCanons:        canonMapFromGame(g.CanonDeck),
		deployAreas:         das,
		actionButton:        ab,
		deployCannonTrigger: deployCannonTrigger,
		moveCannonTrigger:   moveCannonTrigger,
	}
}

func canonMapFromGame(gcd game.CanonDeck) map[int]*ebitenCanon {
	cs := make(map[int]*ebitenCanon, gcd.CanonCapacity())
	for formationPlacement := 0; formationPlacement < gcd.CanonCapacity(); formationPlacement++ {
		canon := gcd.Canons[game.BattleGroundColumn(formationPlacement)]
		if canon != nil {
			println(fmt.Sprintf("Setting cannon to position %d", formationPlacement))
			ec := newEbitenCanon(
				*canon,
				formationPlacement,
				LoadedImages[RegularCanon],
				getCanonCenterX(formationPlacement, gcd.CanonCapacity()),
				canonYPlacement,
			)
			cs[formationPlacement] = &ec
		} else {
			println(fmt.Sprintf("No cannon to position %d", formationPlacement))
		}
	}
	return cs
}

func getCanonCenterX(formationPlacement, numberCanons int) float64 {
	availableWidth := float64(constants.ScreenWidth / numberCanons)
	return availableWidth*float64(formationPlacement) + availableWidth/2
}

func (ecd *EbitenCanonDeck) FireCanons(g *game.CanonTDGame) {
	ecd.ebitenCanons = canonMapFromGame(g.CanonDeck)
	for _, canon := range ecd.ebitenCanons {
		if canon != nil {
			canon.fire()
		}
	}
}

func (ecd *EbitenCanonDeck) Draw(screen *ebiten.Image) {
	for _, canon := range ecd.ebitenCanons {
		canon.draw(screen)
	}
	for _, deployArea := range ecd.deployAreas {
		deployArea.draw(screen)
	}
	ecd.actionButton.draw(screen)
}

func (ecd *EbitenCanonDeck) Update() {
	for _, canon := range ecd.ebitenCanons {
		canon.update()
	}

	ecd.actionButton.update()

	for _, deployArea := range ecd.deployAreas {
		deployArea.update(ecd.draggedSprite)
	}
}

func (ecd *EbitenCanonDeck) InitDrag() {
	ecd.actionButton.canonSprite.InitDrag()
	if ecd.actionButton.canonSprite.IsDragged {
		ecd.draggedSprite = ecd.actionButton.canonSprite
		return
	}
	for _, canon := range ecd.ebitenCanons {
		canon.sprite.InitDrag()
		if canon.sprite.IsDragged {
			ecd.draggedSprite = canon.sprite
			return
		}
	}
}

func (ecd *EbitenCanonDeck) ReleaseDrag() {
	if ecd.draggedSprite != nil {
		if ecd.actionButton.canonSprite.IsDragged {
			ecd.deploy()
			ecd.actionButton.canonSprite.ReleaseDrag()
		}
		for _, canon := range ecd.ebitenCanons {
			if canon.sprite.IsDragged {
				ecd.moveCanon(canon)
				canon.sprite.ReleaseDrag()
			}
		}
		ecd.draggedSprite = nil
	}
}

func (ecd *EbitenCanonDeck) moveCanon(canon *ebitenCanon) {
	deployedArea := ecd.getDeployedAreaPosition()
	if deployedArea != nil {
		formationPlacement := *deployedArea
		if canon.formationPlacement == formationPlacement {
			return
		}
		ecd.moveCannonTrigger(canon.formationPlacement, formationPlacement)
	}
}

func (ecd *EbitenCanonDeck) deploy() {
	deployedArea := ecd.getDeployedAreaPosition()
	if deployedArea != nil {
		ecd.deployCannonTrigger(*deployedArea)
	}
}

func (ecd *EbitenCanonDeck) getDeployedAreaPosition() *int {
	for position, da := range ecd.deployAreas {
		if ebiten_sprite.Collision(da, ecd.draggedSprite) {
			return &position
		}
	}
	return nil
}

func (ecd *EbitenCanonDeck) CurrentBullets() []*EbitenCanonBullet {
	bullets := make([]*EbitenCanonBullet, 0, len(ecd.ebitenCanons))
	for _, ec := range ecd.ebitenCanons {
		if ec.bullet != nil {
			bullets = append(bullets, ec.bullet)
		}
	}
	return bullets
}
