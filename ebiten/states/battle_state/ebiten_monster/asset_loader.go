package ebiten_monster

import (
	"bytes"
	"canon-tower-defense/ebiten/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

var LoadedImages map[string]*ebiten.Image

const BeholderIdle string = "beholder"
const BeholderHit string = "beholderHit"
const BeholderDead string = "BeholderDead"

func LoadBattleImages() {
	// Initialize the global image map
	LoadedImages = make(map[string]*ebiten.Image)

	loadImage(BeholderIdle, assets.Beholder)
	loadImage(BeholderHit, assets.BeholderHit)
	loadImage(BeholderDead, assets.BeholderDie)
}

func loadImage(name string, data []byte) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	LoadedImages[name] = ebiten.NewImageFromImage(img)
}
