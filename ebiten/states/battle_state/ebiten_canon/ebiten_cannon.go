package ebiten_canon

import (
	"bytes"
	"canon-tower-defense/ebiten/assets"
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"image"
	"log"
	"strconv"
)

const BulletSpeed float64 = 5

type ebitenCanon struct {
	sprite             *ebiten_sprite.EbitenDraggableSprite
	Canon              game.Canon
	formationPlacement int
	canonPlacedImage   *ebiten.Image
	bulletImage        *ebiten.Image
	bullet             *EbitenCanonBullet
}

func newEbitenCanon(
	canon game.Canon,
	formationPlacement int,
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
		sprite:             ebiten_sprite.NewFromSprite(sprite),
		Canon:              canon,
		formationPlacement: formationPlacement,
		canonPlacedImage:   cImage,
		bulletImage:        bulletImage,
		bullet:             nil,
	}
}

func (ec *ebitenCanon) fire() {
	// passing the bullet speed because it kind of belongs to the cannon type, etc
	bullet := NewBullet(ec, ec.bulletImage, BulletSpeed, ec.sprite.PosX, CanonYPlacement)
	ec.bullet = &bullet
}

func (ec *ebitenCanon) update() {
	if ec.bullet != nil {
		ec.bullet.update()
		if !ebiten_sprite.SpriteInScreen(ec.bullet.BulletSprite) {
			ec.bullet = nil
		}
	}
	ec.sprite.Update()
}

func (ec *ebitenCanon) draw(screen *ebiten.Image) {
	ec.sprite.Draw(screen)
	if ec.bullet != nil {
		ec.bullet.draw(screen, int(ec.Canon.Damage))
	}
	ec.drawDamage(screen)
}

func (ec *ebitenCanon) drawDamage(screen *ebiten.Image) {
	basicFont := basicfont.Face7x13
	imgWidth := float64(ec.sprite.Image.Bounds().Dx())
	imgHeight := float64(ec.sprite.Image.Bounds().Dy())

	posX := ec.sprite.PosX + (imgWidth / 2) - 4
	posY := ec.sprite.PosY + imgHeight + 25

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1.5, 1.5)
	op.GeoM.Translate(posX, posY)

	text.DrawWithOptions(screen,
		strconv.FormatInt(int64(ec.Canon.Damage), 10),
		basicFont,
		op,
	)
}
