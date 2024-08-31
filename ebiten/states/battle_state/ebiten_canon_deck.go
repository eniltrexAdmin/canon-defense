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
	ebitenCanons []ebitenCanon
	actionButton ebitenActionButton
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

	var cs []ebitenCanon
	for j, canon := range cd.Canons {
		ec := newEbitenCanon(canon, canonImage, j, availableWidth)
		cs = append(cs, ec)
	}

	ab := newEbitenActionButton(canonImage, pedestalImage, constants.ScreenWidth)

	return ebitenCanonDeck{
		ebitenCanons: cs,
		actionButton: ab,
	}
}

func (ecd ebitenCanonDeck) draw(screen *ebiten.Image) {
	for _, canon := range ecd.ebitenCanons {
		canon.draw(screen)
	}
	ecd.actionButton.draw(screen)
}

func (ecd ebitenCanonDeck) click(x, y int) *ebitenCanon {
	for _, ec := range ecd.ebitenCanons {
		if ec.inBounds(float32(x), float32(y)) {
			return &ec
		}
	}
	return nil
}
