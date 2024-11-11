package battle_state

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"math"
)

type ebitenCanonBullet struct {
	bulletSprite ebiten_sprite.EbitenSprite
	bulletSpeed  float64
}

func NewBullet(bulletImage *ebiten.Image, speed float64, posX, posY float64) ebitenCanonBullet {

	imgWidth := bulletImage.Bounds().Dx()
	imgHeight := bulletImage.Bounds().Dy()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(90 * math.Pi / 180)
	op.GeoM.Translate(posX+24, posY-20)

	tx, ty := op.GeoM.Element(0, 2), op.GeoM.Element(1, 2)

	log.Printf("Transformed Top-Left X: %f, Y: %f", tx, ty)

	sprite := ebiten_sprite.NewEbitenSprite(
		tx-float64(imgHeight),
		ty,
		float64(imgHeight),
		float64(imgWidth), //reversing since in draw I am rotating it
		bulletImage,
		1,
	)

	return ebitenCanonBullet{bulletSprite: sprite, bulletSpeed: speed}
}

func (eb *ebitenCanonBullet) update() {
	eb.bulletSprite.PosY -= eb.bulletSpeed
}

func (eb *ebitenCanonBullet) draw(screen *ebiten.Image, damage int) {
	imgHeight := float64(eb.bulletSprite.Image.Bounds().Dy()) // it's rotated
	scale := 1 + float64(damage)/10

	realHeight := imgHeight * scale

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(90 * math.Pi / 180)
	op.GeoM.Scale(scale, scale)
	// it's rotated so the height is used for centering width
	op.GeoM.Translate(eb.bulletSprite.PosX+realHeight, eb.bulletSprite.PosY)
	screen.DrawImage(eb.bulletSprite.Image, op)
}
