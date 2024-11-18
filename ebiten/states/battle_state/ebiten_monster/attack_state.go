package ebiten_monster

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

const AttackStateName = "Attacking!"
const MovementSpeed float64 = 2

type MonsterAttackState struct {
	sprite      *ebiten_sprite.EbitenAnimatedSprite
	context     *EbitenMonster
	destination ebiten_sprite.ScreenCoordinate
}

func newMonsterAttackState(currentScreenPosition ebiten_sprite.ScreenCoordinate, stepSize float64) *MonsterAttackState {
	beholderAttack := ebiten_sprite.NewFromCentralPoint(
		currentScreenPosition.X,
		currentScreenPosition.Y,
		LoadedImages[BeholderAttack],
		128,
		128,
		1,
		0.1)

	// TODO thats probably should be given as parameter
	// especially if sometimes moving horizontally.
	d := ebiten_sprite.ScreenCoordinate{
		X: currentScreenPosition.X,
		Y: stepSize + currentScreenPosition.Y,
	}

	return &MonsterAttackState{
		sprite:      &beholderAttack,
		destination: d,
	}
}

func (m *MonsterAttackState) setContext(context *EbitenMonster) {
	m.context = context
	m.context.LifeLine.Hide()
}

func (m *MonsterAttackState) draw(screen *ebiten.Image) {
	m.sprite.Draw(screen)
}

func (m *MonsterAttackState) update() {
	m.sprite.Move(m.destination, MovementSpeed)
	m.sprite.Update()
	if m.sprite.Position() == m.destination {
		idleState := newIdleMonster(m.Coordinates())
		m.context.transitionTo(&idleState)
	}
}

func (m *MonsterAttackState) stateName() string {
	return AttackStateName
}

func (m *MonsterAttackState) Coordinates() ebiten_sprite.ScreenCoordinate {
	return m.sprite.Position()
}

func (m *MonsterAttackState) GetRectangle() image.Rectangle {
	return m.sprite.GetRectangle()
}
