package battle_state

import (
	"bytes"
	"canon-tower-defense/ebiten/assets"
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
	"math"
)

// TODO move functionality to "animated Sprite"

const animationSpeed float64 = 0.15

type ebitenMonster struct {
	monster *game.Monster
	sprite  ebiten_sprite.EbitenSprite
	frame   float64
}

func (m *ebitenMonster) draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(-float64(200)/2, -float64(200)/2-10)
	op.GeoM.Translate(m.sprite.PosX, m.sprite.PosY)
	i := int(math.Floor(m.frame))
	sx, sy := i*100, 0

	screen.DrawImage(m.sprite.Image.SubImage(
		image.Rect(i*100, sy, sx+100, sy+100)).(*ebiten.Image), op)
}

func (m *ebitenMonster) update() {
	m.frame += animationSpeed
	if m.frame >= 6 {
		m.frame = 0
	}
}

func NewEbitenMonster(monster *game.Monster, posX, posY float64) ebitenMonster {
	// switch asset by name
	img, _, err := image.Decode(bytes.NewReader(assets.Skeleton))
	if err != nil {
		log.Fatal(err)
	}
	skeletonImage := ebiten.NewImageFromImage(img)

	sprite := ebiten_sprite.NewEbitenSprite(posX, posY, TileSize-5, TileSize-5, skeletonImage, 1)
	return ebitenMonster{
		sprite:  sprite,
		monster: monster,
	}
}
