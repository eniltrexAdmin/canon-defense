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
	monster   *game.Monster
	sprite    ebiten_sprite.EbitenSprite
	hurtImage *ebiten.Image
	frame     float64
	hit       bool
}

func (m *ebitenMonster) draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2.5, 2.5)
	op.GeoM.Translate(-float64(250)/2, -float64(250)/2-10)
	op.GeoM.Translate(m.sprite.PosX, m.sprite.PosY)
	i := int(math.Floor(m.frame))
	sx, sy := i*100, 0

	if m.hit {
		screen.DrawImage(m.hurtImage.SubImage(
			image.Rect(i*100, sy, sx+100, sy+100)).(*ebiten.Image), op)
	} else {
		screen.DrawImage(m.sprite.Image.SubImage(
			image.Rect(i*100, sy, sx+100, sy+100)).(*ebiten.Image), op)
	}
}

func (m *ebitenMonster) update(bullets []*ebitenCanonBullet) {
	m.frame += animationSpeed
	if m.frame >= 6 {
		m.frame = 0
	}

	for _, bullet := range bullets {
		if m.sprite.Collision(bullet.bulletSprite) {
			m.hit = true
		} else {
			m.hit = false // TODO that's going to be a problem.
		}
	}
	if m.hit {
		if m.frame >= 4 {
			m.frame = 0
		}
	}
}

func NewEbitenMonster(monster *game.Monster, posX, posY float64) ebitenMonster {
	// switch asset by name
	img, _, err := image.Decode(bytes.NewReader(assets.Skeleton))
	if err != nil {
		log.Fatal(err)
	}
	skeletonImage := ebiten.NewImageFromImage(img)

	img2, _, err := image.Decode(bytes.NewReader(assets.SkeletonHurt))
	if err != nil {
		log.Fatal(err)
	}
	skeletonHurt := ebiten.NewImageFromImage(img2)

	sprite := ebiten_sprite.NewEbitenSprite(posX, posY, TileSize-5, TileSize-5, skeletonImage, 1)
	return ebitenMonster{
		sprite:    sprite,
		hurtImage: skeletonHurt,
		monster:   monster,
		hit:       false,
	}
}
