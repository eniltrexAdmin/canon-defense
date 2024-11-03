package battle_state

import (
	"bytes"
	"canon-tower-defense/ebiten/assets"
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

const canonYPlacement float64 = 550

type ebitenCanonDeck struct {
	ebitenCanons  map[int]*ebitenCanon
	deployAreas   map[int]*deployArea
	actionButton  ebitenActionButton
	gameCanonDeck *game.CanonDeck
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

	availableWidth := float64(constants.ScreenWidth / len(cd.Canons))

	cs := make(map[int]*ebitenCanon, len(cd.Canons))
	das := make(map[int]*deployArea, len(cd.Canons))

	color := ebiten_sprite.RandomColor()
	for formationPlacement, canon := range cd.Canons {
		centerX := getCanonCenterX(formationPlacement, len(cd.Canons))
		if canon != nil {
			ec := newEbitenCanon(
				*canon,
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
	ecd.actionButton.update(ecd)
	for _, canon := range ecd.ebitenCanons {
		canon.update()
	}
	for _, deployArea := range ecd.deployAreas {
		dragged := ecd.actionButton.dragged
		draggedSprite := ecd.actionButton.canonSprite
		deployArea.update(dragged, draggedSprite)
	}
}

func (ecd *ebitenCanonDeck) deploy(canonSprite ebiten_sprite.EbitenSprite) {
	for position, da := range ecd.deployAreas {
		if ebiten_sprite.Collision(da, canonSprite) {
			// TODO this build of cannon will definitely have more domain.
			// probably both this and the "ebiten" version.

			canon := game.BuildCanon(1)
			ecd.gameCanonDeck.DeployCannon(game.BattleGroundColumn(position), &canon)

			ec := newEbitenCanon(
				*ecd.gameCanonDeck.Canons[position],
				canonSprite.Image,
				getCanonCenterX(position, len(ecd.gameCanonDeck.Canons)),
				canonYPlacement,
			)
			ecd.ebitenCanons[position] = &ec
			ec.fire()
		}
	}
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
