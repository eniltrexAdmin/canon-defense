package battle_state

import (
	"bytes"
	"canon-tower-defense/ebiten/assets"
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

const BulletSpeed float64 = 2.5

type ebitenCanon struct {
	sprite           *ebiten_sprite.EbitenSprite
	canon            game.Canon
	canonPlacedImage *ebiten.Image
	bulletImage      *ebiten.Image
	bullet           *ebitenCanonBullet
}

func newEbitenCanon(
	canon game.Canon,
	cImage *ebiten.Image,
	centerX float64,
	centerY float64,
) ebitenCanon {

	sprite := ebiten_sprite.NewStaticFromCentralPoint(
		centerX,
		centerY,
		cImage,
	)

	// TODO assets management should eventually be centralized.
	img3, _, err := image.Decode(bytes.NewReader(assets.Bullet))
	if err != nil {
		log.Fatal(err)
	}
	bulletImage := ebiten.NewImageFromImage(img3)

	return ebitenCanon{
		sprite:           &sprite,
		canon:            canon,
		canonPlacedImage: cImage,
		bulletImage:      bulletImage,
		bullet:           nil,
	}
}

func (ec *ebitenCanon) fire() {
	bullet := NewBullet(ec.bulletImage, BulletSpeed, ec.sprite.PosX, canonYPlacement)
	ec.bullet = &bullet
}

func (ec *ebitenCanon) update() {
	if ec.bullet != nil {
		ec.bullet.update()
		if ec.bullet.bulletSprite.PosY < 0 {
			ec.bullet = nil
		}
	}
}

func (ec *ebitenCanon) draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ec.sprite.PosX, ec.sprite.PosY)
	screen.DrawImage(ec.canonPlacedImage, op)

	if ec.bullet != nil {
		ec.bullet.draw(screen)
	}
}
