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

const animationSpeed float64 = 0.15

type ebitenMonster struct {
	monster       *game.Monster
	sprite        ebiten_sprite.EbitenSprite
	hurtImage     *ebiten.Image
	frame         float64
	currentSprite *ebiten_sprite.EbitenAnimatedSprite
	idleSprite    ebiten_sprite.EbitenAnimatedSprite
	hitSprite     ebiten_sprite.EbitenAnimatedSprite
}

func (m *ebitenMonster) draw(screen *ebiten.Image) {
	m.currentSprite.Draw(screen)
	//m.currentSprite.DrawDebug(screen)
}

func (m *ebitenMonster) update(bullets []*ebitenCanonBullet) {
	m.currentSprite = &m.idleSprite
	if m.IsHit(bullets) {
		m.currentSprite = &m.hitSprite
	}
	m.currentSprite.Update()
}

func (m *ebitenMonster) IsHit(bullets []*ebitenCanonBullet) bool {
	for _, bullet := range bullets {
		if ebiten_sprite.Collision(m.currentSprite, bullet.bulletSprite) {
			return true
		}
	}
	return false
}

func NewEbitenMonster(monster *game.Monster, posX, posY float64) ebitenMonster {
	// switch asset by name
	img, _, err := image.Decode(bytes.NewReader(assets.Beholder))
	if err != nil {
		log.Fatal(err)
	}
	beholderImage := ebiten.NewImageFromImage(img)

	img2, _, err := image.Decode(bytes.NewReader(assets.BeholderHit))
	if err != nil {
		log.Fatal(err)
	}
	bhImage := ebiten.NewImageFromImage(img2)

	beholder := ebiten_sprite.NewFromCentralPoint(
		posX,
		posY,
		beholderImage,
		64,
		64,
		1,
		0.1)

	beholderHit := ebiten_sprite.NewFromCentralPoint(
		posX,
		posY,
		bhImage,
		64,
		64,
		1,
		0.1)

	return ebitenMonster{
		monster:       monster,
		idleSprite:    beholder,
		hitSprite:     beholderHit,
		currentSprite: &beholder,
	}
}
