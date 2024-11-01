package battle_state

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/game"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

const canonYPlacement float64 = 500

const BulletSpeed float64 = 1

// TODO this is not good. Bullet and Cannon are probably independent.
// as well as "DeployArea".

type ebitenCanon struct {
	ebiten_sprite.EbitenSprite
	formationPlacement int
	canon              *game.Canon
	canonPlacedImage   *ebiten.Image
	bulletImage        *ebiten.Image
	bullet             *ebitenCanonBullet
}

func newEbitenCanon(canon *game.Canon, cImage *ebiten.Image, bulletImage *ebiten.Image, formationPlacement int, availableWidth int) ebitenCanon {

	imgHeight := cImage.Bounds().Dy()
	imgWidth := cImage.Bounds().Dx()

	centerSpace := availableWidth / 2
	imgStartingXPoint := centerSpace - imgWidth/2

	fmt.Println(fmt.Sprintf("placing canon in: %.2f",
		float32((availableWidth*formationPlacement)+imgStartingXPoint)),
	)

	sprite := ebiten_sprite.NewEbitenSprite(
		float64((availableWidth*formationPlacement)+imgStartingXPoint),
		canonYPlacement,
		float64(imgWidth),
		float64(imgHeight),
		cImage,
		1,
	)

	return ebitenCanon{
		EbitenSprite:       sprite,
		formationPlacement: formationPlacement,
		canon:              canon,
		canonPlacedImage:   cImage,
		bulletImage:        bulletImage,
		bullet:             nil,
	}
}

func (ec *ebitenCanon) placeCannon(cannon *game.Canon) {
	ec.canon = cannon
	bullet := NewBullet(ec.bulletImage, BulletSpeed, ec.EbitenSprite.PosX, canonYPlacement)
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

	var fill float32 = 0.5
	if ec.canon != nil {
		fill = 1
	}
	op.ColorScale.ScaleAlpha(fill)
	op.GeoM.Translate(ec.PosX, ec.PosY)
	screen.DrawImage(ec.canonPlacedImage, op)

	if ec.bullet != nil {
		ec.bullet.draw(screen)
	}
}
