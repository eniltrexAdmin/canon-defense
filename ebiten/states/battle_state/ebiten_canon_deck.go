package battle_state

import (
	"bytes"
	"canon-tower-defense/ebiten/assets"
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

type ebitenCanonDeck struct {
	ebitenCanons  []*ebitenCanon
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

	availableWidth := constants.ScreenWidth / len(cd.Canons)

	img3, _, err := image.Decode(bytes.NewReader(assets.Bullet))
	if err != nil {
		log.Fatal(err)
	}
	bulletImage := ebiten.NewImageFromImage(img3)

	var cs []*ebitenCanon
	for j, canon := range cd.Canons {
		ec := newEbitenCanon(canon, canonImage, bulletImage, j, availableWidth)
		cs = append(cs, &ec)
	}

	ab := newEbitenActionButton(canonImage, pedestalImage, constants.ScreenWidth)

	return ebitenCanonDeck{
		ebitenCanons:  cs,
		actionButton:  ab,
		gameCanonDeck: &cd,
	}
}

func (ecd *ebitenCanonDeck) draw(screen *ebiten.Image) {
	for _, canon := range ecd.ebitenCanons {
		canon.draw(screen)
	}
	ecd.actionButton.draw(screen)
}

func (ecd *ebitenCanonDeck) update() {
	for _, canon := range ecd.ebitenCanons {
		canon.update()
	}
}

func (ecd *ebitenCanonDeck) deploy(x, y int) { // TODO we pass the "cannon" here, we deploy a cannon.)
	for position, ec := range ecd.ebitenCanons {
		if ec.InBounds(x, y) {
			canon := game.BuildCanon(1)
			ecd.gameCanonDeck.DeployCannon(game.BattleGroundColumn(position), &canon)
			ecd.ebitenCanons[position].placeCannon(ecd.gameCanonDeck.Canons[position])
		}
	}
}
