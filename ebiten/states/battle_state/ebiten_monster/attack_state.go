package ebiten_monster

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"github.com/hajimehoshi/ebiten/v2"
)

type MonsterAttackState struct {
	sprite      *ebiten_sprite.EbitenAnimatedSprite
	context     *EbitenMonster
	advanceStep float64
}

func newMonsterAttackState(mis *MonsterIdleState) *MonsterAttackState {
	coordinate := mis.sprite.Position()

	beholderHit := ebiten_sprite.NewFromCentralPoint(
		coordinate.X,
		coordinate.Y,
		LoadedImages[BeholderAttack],
		64,
		64,
		1,
		0.1)

	return &MonsterAttackState{
		sprite: &beholderHit,
	}
}

func (m *MonsterAttackState) setContext(context *EbitenMonster) {
	m.context = context
}

func (m *MonsterAttackState) draw(screen *ebiten.Image) {
	m.sprite.Draw(screen)
}

func (m *MonsterAttackState) update() {
	m.sprite.Update()
}

func (m *MonsterAttackState) stateName() string {
	return "attack State"
}
