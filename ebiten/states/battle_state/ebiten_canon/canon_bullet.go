package ebiten_canon

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"math"
)

type EbitenCanonBullet struct {
	BulletSprite ebiten_sprite.EbitenSprite
	bulletSpeed  float64
	Canon        *ebitenCanon
}

func NewBullet(canon *ebitenCanon, bulletImage *ebiten.Image, speed float64, posX, posY float64) EbitenCanonBullet {

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

	return EbitenCanonBullet{BulletSprite: sprite, bulletSpeed: speed, Canon: canon}
}

func (eb *EbitenCanonBullet) update() {
	eb.BulletSprite.PosY -= eb.bulletSpeed
}

func (eb *EbitenCanonBullet) draw(screen *ebiten.Image, damage int) {
	imgHeight := float64(eb.BulletSprite.Image.Bounds().Dy()) // it's rotated
	scale := 1 + float64(damage)/10

	realHeight := imgHeight * scale

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(90 * math.Pi / 180)
	op.GeoM.Scale(scale, scale)
	// it's rotated so the height is used for centering width
	op.GeoM.Translate(eb.BulletSprite.PosX+realHeight, eb.BulletSprite.PosY)
	screen.DrawImage(eb.BulletSprite.Image, op)
}
