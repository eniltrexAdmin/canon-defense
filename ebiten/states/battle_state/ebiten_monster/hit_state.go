package ebiten_monster

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_canon"
	"github.com/hajimehoshi/ebiten/v2"
)

type MonsterHitState struct {
	sprite        *ebiten_sprite.EbitenAnimatedSprite
	context       *EbitenMonster
	hittingBullet *ebiten_canon.EbitenCanonBullet
}

func newMonsterHitState(mis *MonsterIdleState, bullet *ebiten_canon.EbitenCanonBullet) *MonsterHitState {
	coordinate := mis.sprite.Position()

	beholderHit := ebiten_sprite.NewFromCentralPoint(
		coordinate.X,
		coordinate.Y,
		LoadedImages[BeholderHit],
		64,
		64,
		1.1,
		0.1)

	return &MonsterHitState{
		sprite:        &beholderHit,
		hittingBullet: bullet,
	}
}

func (m *MonsterHitState) setContext(context *EbitenMonster) {
	m.context = context
	// whenever this status is set, game API interface is called
	context.game.HitMonster(&m.hittingBullet.Canon.Canon, m.context.monster)
	context.LifeLine.SetCurrentLife(int(context.monster.HealthPoints))
}

func (m *MonsterHitState) draw(screen *ebiten.Image) {
	m.sprite.Draw(screen)
}

func (m *MonsterHitState) update() {
	m.sprite.Update()
}

func (m *MonsterHitState) updateDeckFiring(bullets []*ebiten_canon.EbitenCanonBullet) {
	if !ebiten_sprite.Collision(m.sprite, m.hittingBullet.BulletSprite) {
		if m.context.monster.IsAlive() {
			idleState := newFromHitState(m)
			m.context.transitionTo(&idleState)
		} else {
			// go to dead state
			deadState := newDeadFromHitState(m)
			m.context.transitionTo(&deadState)
		}
	}

	m.sprite.Update()
}

func (m *MonsterHitState) updateAttack() {
	m.sprite.Update()
}

func (m *MonsterHitState) stateName() string {
	return "hit State"
}
