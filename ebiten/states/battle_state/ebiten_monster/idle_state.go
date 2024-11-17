package ebiten_monster

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_canon"
	"github.com/hajimehoshi/ebiten/v2"
)

type MonsterIdleState struct {
	sprite  *ebiten_sprite.EbitenAnimatedSprite
	context *EbitenMonster
}

func newIdleMonster(posX, posY float64) MonsterIdleState {
	beholder := ebiten_sprite.NewFromCentralPoint(
		posX,
		posY,
		LoadedImages[BeholderIdle],
		64,
		64,
		1,
		0.1)
	return MonsterIdleState{
		sprite: &beholder,
	}
}

func newFromHitState(hs *MonsterHitState) MonsterIdleState {
	c := hs.sprite.Position()
	return newIdleMonster(c.X, c.Y)
}

func (m *MonsterIdleState) setContext(context *EbitenMonster) {
	m.context = context
}

func (m *MonsterIdleState) draw(screen *ebiten.Image) {
	m.sprite.Draw(screen)
}

func (m *MonsterIdleState) update() {
	m.sprite.Update()
}

func (m *MonsterIdleState) updateDeckFiring(bullets []*ebiten_canon.EbitenCanonBullet) {
	for _, bullet := range bullets {
		if ebiten_sprite.Collision(m.sprite, bullet.BulletSprite) {
			hitState := newMonsterHitState(m, bullet)
			m.context.transitionTo(hitState)
		}
	}

	m.sprite.Update()
}

func (m *MonsterIdleState) updateAttack() {
	m.sprite.Update()
}

func (m *MonsterIdleState) stateName() string {
	return "Idle State"
}
