package battle_state

import (
	"bytes"
	"canon-tower-defense/ebiten/assets"
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
	"image"
	"image/color"
	"log"
	"strconv"
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
	ec.sprite.Draw(screen)
	if ec.bullet != nil {
		ec.bullet.draw(screen)
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
		strconv.FormatInt(int64(ec.canon.Damage), 10),
		basicFont,
		op,
	)
	gray := color.RGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xff}
	// TODO get size of text for better centering of circle
	vector.StrokeCircle(screen, float32(posX)+4, float32(posY)-6,
		13, float32(1), gray, true)
}
