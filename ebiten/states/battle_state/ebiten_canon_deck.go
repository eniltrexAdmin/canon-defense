package battle_state

import (
	"bytes"
	"canon-tower-defense/assets"
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

type ebitenCanonDeck struct {
	ebitenCanons []ebitenCanon
}

func newEbitenCanonDeck(cd game.CanonDeck) ebitenCanonDeck {
	img, _, err := image.Decode(bytes.NewReader(assets.RegularCanon))
	if err != nil {
		log.Fatal(err)
	}
	canonImage := ebiten.NewImageFromImage(img)
	availableWidth := constants.ScreenWidth / len(cd.Canons)

	var cs []ebitenCanon
	for j, canon := range cd.Canons {
		ec := newEbitenCanon(canon, canonImage, j, availableWidth)
		cs = append(cs, ec)
	}
	return ebitenCanonDeck{ebitenCanons: cs}
}

func (ecd ebitenCanonDeck) draw(screen *ebiten.Image) {
	for _, canon := range ecd.ebitenCanons {
		canon.draw(screen)
	}
}

func (ecd ebitenCanonDeck) click(x, y int) *ebitenCanon {
	for _, ec := range ecd.ebitenCanons {
		if ec.inBounds(float32(x), float32(y)) {
			return &ec
		}
	}
	return nil
}
