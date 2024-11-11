package battle_state

import (
	"bytes"
	"canon-tower-defense/ebiten/assets"
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/game"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
	"log"
)

const canonYPlacement float64 = 550

type ebitenCanonDeck struct {
	ebitenCanons  map[int]*ebitenCanon
	deployAreas   map[int]*deployArea
	actionButton  ebitenActionButton
	gameCanonDeck *game.CanonDeck
	Firing        bool
}

func newEbitenCanonDeck(cd game.CanonDeck) ebitenCanonDeck {
	img, _, err := image.Decode(bytes.NewReader(assets.RegularCanon))
	if err != nil {
		log.Fatal(err)
	}
	canonImage := ebiten.NewImageFromImage(img)

	img2, _, err := image.Decode(bytes.NewReader(assets.CanonPedestal))
	if err != nil {
		log.Fatal(err)
	}
	pedestalImage := ebiten.NewImageFromImage(img2)

	availableWidth := float64(constants.ScreenWidth / cd.CanonCapacity())

	cs := make(map[int]*ebitenCanon, cd.CanonCapacity())
	das := make(map[int]*deployArea, cd.CanonCapacity())

	color := ebiten_sprite.RandomColor()
	for formationPlacement := 0; formationPlacement < cd.CanonCapacity(); formationPlacement++ {
		centerX := getCanonCenterX(formationPlacement, cd.CanonCapacity())
		canon := cd.Canons[game.BattleGroundColumn(formationPlacement)]
		if canon != nil {
			ec := newEbitenCanon(
				*canon,
				formationPlacement,
				canonImage,
				centerX,
				canonYPlacement,
			)
			cs[formationPlacement] = &ec
		}

		da := NewDeployAreaFromCentralPoint(
			centerX, canonYPlacement,
			availableWidth-20, 60,
			color,
		)
		das[formationPlacement] = &da
	}

	ab := newEbitenActionButton(canonImage, pedestalImage, constants.ScreenWidth)

	return ebitenCanonDeck{
		ebitenCanons:  cs,
		deployAreas:   das,
		actionButton:  ab,
		gameCanonDeck: &cd,
		Firing:        false,
	}
}

func getCanonCenterX(formationPlacement, numberCanons int) float64 {
	availableWidth := float64(constants.ScreenWidth / numberCanons)
	return availableWidth*float64(formationPlacement) + availableWidth/2
}

func (ecd *ebitenCanonDeck) draw(screen *ebiten.Image) {
	for _, canon := range ecd.ebitenCanons {
		canon.draw(screen)
	}
	for _, deployArea := range ecd.deployAreas {
		deployArea.draw(screen)
	}
	ecd.actionButton.draw(screen)
}

func (ecd *ebitenCanonDeck) update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		ecd.initDrag()
	}
	for _, canon := range ecd.ebitenCanons {
		canon.update(ecd)
	}

	ecd.actionButton.update(ecd)
	for _, deployArea := range ecd.deployAreas {
		deployArea.update(ecd.draggedSprite())
	}
}

func (ecd *ebitenCanonDeck) firingUpdate() {
	bulletsInField := false
	for _, canon := range ecd.ebitenCanons {
		canon.firingUpdate()
		if canon.bullet != nil {
			bulletsInField = true
		}
	}
	if bulletsInField == false {
		ecd.Firing = false
	}
}

func (ecd *ebitenCanonDeck) initDrag() {
	x, y := ebiten.CursorPosition()
	ecd.actionButton.initDrag(x, y)
	for _, canon := range ecd.ebitenCanons {
		canon.initDrag(x, y)
	}
}

func (ecd *ebitenCanonDeck) draggedSprite() *ebiten_sprite.EbitenSprite {
	if ecd.actionButton.dragged {
		return &ecd.actionButton.canonSprite
	}
	for _, canon := range ecd.ebitenCanons {
		if canon.dragged {
			return canon.sprite
		}
	}
	return nil
}

func (ecd *ebitenCanonDeck) moveCanon(canon *ebitenCanon) {
	deployedArea := ecd.getDeployedAreaPosition(*canon.sprite)
	if deployedArea != nil {
		formationPlacement := *deployedArea
		if canon.formationPlacement == formationPlacement {
			return
		}
		ecd.gameCanonDeck.MoveCanon(
			game.BattleGroundColumn(canon.formationPlacement),
			game.BattleGroundColumn(formationPlacement))

		ecd.finishTurn(canon.sprite)
	}
}

func (ecd *ebitenCanonDeck) deploy(canonSprite ebiten_sprite.EbitenSprite) {
	deployedArea := ecd.getDeployedAreaPosition(canonSprite)
	if deployedArea != nil {
		formationPlacement := *deployedArea
		// TODO this build of cannon will definitely have more domain.
		// probably both this and the "ebiten" version.

		// This is a very crucial part, actually. The backend game knows whether to merge
		// or create a canon
		canon := game.BuildCanon(1)
		ecd.gameCanonDeck.DeployCannon(game.BattleGroundColumn(formationPlacement), &canon)

		ecd.finishTurn(&canonSprite)
	}
}

func (ecd *ebitenCanonDeck) finishTurn(draggedSprite *ebiten_sprite.EbitenSprite) {
	for formationPlacement := 0; formationPlacement < ecd.gameCanonDeck.CanonCapacity(); formationPlacement++ {
		canon := ecd.gameCanonDeck.Canons[game.BattleGroundColumn(formationPlacement)]
		if canon != nil {
			println(fmt.Sprintf("Setting cannon to position %d", formationPlacement))
			ec := newEbitenCanon(
				*canon,
				formationPlacement,
				draggedSprite.Image,
				getCanonCenterX(formationPlacement, ecd.gameCanonDeck.CanonCapacity()),
				canonYPlacement,
			)
			ecd.ebitenCanons[formationPlacement] = &ec
			ec.fire()
		} else {
			println(fmt.Sprintf("Deleting cannon to position %d", formationPlacement))
			delete(ecd.ebitenCanons, formationPlacement)
		}
	}
	ecd.Firing = true
}

func (ecd *ebitenCanonDeck) getDeployedAreaPosition(canonSprite ebiten_sprite.EbitenSprite) *int {
	for position, da := range ecd.deployAreas {
		if ebiten_sprite.Collision(da, canonSprite) {
			return &position
		}
	}
	return nil
}

func (ecd *ebitenCanonDeck) currentBullets() []*ebitenCanonBullet {
	bullets := make([]*ebitenCanonBullet, 0, len(ecd.ebitenCanons))
	for _, ec := range ecd.ebitenCanons {
		if ec.bullet != nil {
			bullets = append(bullets, ec.bullet)
		}
	}
	return bullets
}
