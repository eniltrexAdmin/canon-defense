package ebiten_monster

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_canon"
	"canon-tower-defense/game"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type EbitenMonsterState interface {
	setContext(context *EbitenMonster)
	draw(screen *ebiten.Image)
	update()
	stateName() string
	Coordinates() ebiten_sprite.ScreenCoordinate
	GetRectangle() image.Rectangle
}

// In order to control possible transition between states, in the interface I could add
// hit() so from Idle state I would implement that, but from dead state I would return error "you cannot hit a dead"
// but I believe this is unnecessary, I just implement the actual changes, not all possible combinations.

// Also I have control on the "new" from where I allow the creation, if it accepts on the parameter of a specific state.
// in PHP all states inherit from an abstract and they all have the context, but here, some states have more
// attributes. (which I guess it would be better to have it in the "ebitenmonster?" the only place is the "hitting bullet".

// hitting bullet is available from inside the monster->bullets in field, the bullets in field had to be passed from
// external place. And the change of state is for intenral logic, not like "advance" whcih is a change of state external

// TODO almost for sure that at the end the states will have just teh "context" and the context will have the sprite.
// the states will draw the context.sprite()...

type EbitenMonster struct {
	state          EbitenMonsterState
	monster        *game.Monster
	LifeLine       *LifeLine
	game           *game.CanonTDGame
	bulletsInField []*ebiten_canon.EbitenCanonBullet
}

func NewEbitenMonster(monster *game.Monster, posX, posY float64, game *game.CanonTDGame) *EbitenMonster {
	c := EbitenMonster{
		monster: monster,
		game:    game,
	}
	state := newIdleMonster(ebiten_sprite.ScreenCoordinate{X: posX, Y: posY})

	ll := NewLifeLineFromRectangle(int(monster.MaxLife), state.sprite.GetRectangle())
	c.LifeLine = &ll
	c.transitionTo(&state)
	// If I do not return the pointer, I am returning a copy, which is not the one where "state" is pointing to.
	// so then changing the context, from state.transtitionTo is calling the wrong one, not the one
	// I am returning here, so the one I am returning here cannot be changed of state....
	return &c
}

func (e *EbitenMonster) transitionTo(state EbitenMonsterState) {
	println(fmt.Sprintf("RTransitioning to : %s", state.stateName()))
	e.state = state
	e.state.setContext(e)
}

func (e *EbitenMonster) Advance(stepHeight float64) {
	// when advancing, we clean up the previous turn bullets in field:
	e.bulletsInField = nil
	if !e.monster.IsAlive() {
		return
	}

	destination := stepHeight * float64(e.monster.RowMovement)

	attackState := newMonsterAttackState(e.state.Coordinates(), destination)
	e.transitionTo(attackState)
}

func (e *EbitenMonster) Draw(screen *ebiten.Image) {
	e.state.draw(screen)
	e.LifeLine.Draw(screen)
}

func (e *EbitenMonster) Update() {
	e.state.update()
	e.LifeLine.Update(e.state.GetRectangle())
}

func (e *EbitenMonster) DeckFiring(bullets []*ebiten_canon.EbitenCanonBullet) {
	e.bulletsInField = bullets
}

func (e *EbitenMonster) IsAttacking() bool {
	return e.state.stateName() == AttackStateName
}
