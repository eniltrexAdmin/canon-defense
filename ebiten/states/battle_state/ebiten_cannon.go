package battle_state

import (
	"bytes"
	"canon-tower-defense/ebiten/assets"
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	sprite                               *ebiten_sprite.EbitenSprite
	canon                                game.Canon
	formationPlacement                   int
	canonPlacedImage                     *ebiten.Image
	bulletImage                          *ebiten.Image
	bullet                               *ebitenCanonBullet
	initialPlacementX, initialPlacementY float64
	dragIniX, dragIniY                   float64
	dragged                              bool
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
		sprite:             &sprite,
		canon:              canon,
		formationPlacement: formationPlacement,
		canonPlacedImage:   cImage,
		bulletImage:        bulletImage,
		bullet:             nil,
		initialPlacementX:  sprite.PosX,
		initialPlacementY:  sprite.PosY,
		dragged:            false,
	}
}

func (ec *ebitenCanon) fire() {
	bullet := NewBullet(ec.bulletImage, BulletSpeed, ec.sprite.PosX, canonYPlacement)
	ec.bullet = &bullet
}

// TODO encapsulate drag logic in utils?
func (ec *ebitenCanon) initDrag(x, y int) {
	if ec.sprite.InBounds(x, y) {
		ec.dragIniX = float64(x)
		ec.dragIniY = float64(y)
		ec.dragged = true
	}
}

// TODO this should do other things.
func (ec *ebitenCanon) JustRelease() {
	ec.sprite.PosX = ec.initialPlacementX
	ec.sprite.PosY = ec.initialPlacementY
	ec.dragged = false
}

func (ec *ebitenCanon) update(deck *ebitenCanonDeck) {
	// draw bullet first, even though it should be separated
	if ec.bullet != nil {
		ec.bullet.update()
		if ec.bullet.bulletSprite.PosY < 0 {
			ec.bullet = nil
		}
	}
	if ec.dragged == false {
		return
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		deck.moveCanon(ec)
		ec.JustRelease()
		return
	}

	// TODO encapsulate drag logic in utils?
	// that means here it's still being dragged.
	x, y := ebiten.CursorPosition()

	dragDeltaX := float64(x) - ec.dragIniX + ec.initialPlacementX
	dragDeltaY := float64(y) - ec.dragIniY + ec.initialPlacementY

	ec.sprite.PosX = dragDeltaX
	ec.sprite.PosY = dragDeltaY
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
