package ebiten_canon

import (
	"bytes"
	"canon-tower-defense/ebiten/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

var LoadedImages map[string]*ebiten.Image

const RegularCanon string = "RegularCanon"
const RegularCanonPedestal string = "RegularCanonPedestal"

func LoadBattleImages() {
	// Initialize the global image map
	LoadedImages = make(map[string]*ebiten.Image)

	loadImage(RegularCanon, assets.RegularCanon)
	loadImage(RegularCanonPedestal, assets.CanonPedestal)
}

func loadImage(name string, data []byte) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	LoadedImages[name] = ebiten.NewImageFromImage(img)
}
