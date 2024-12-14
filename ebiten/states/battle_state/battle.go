package battle_state

import (
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/ebiten/states"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_background"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_canon"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_monster"
	"canon-tower-defense/game"
	"canon-tower-defense/game/data"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

const GameOverLine float32 = 510

type BattleState struct {
	game               *game.CanonTDGame
	ebitenCanonDeck    *ebiten_canon.EbitenCanonDeck
	ebitenBattleGround ebiten_background.EbitenBattleGround
	ebitenMonsterTeam  *ebiten_monster.EbitenMonsterTeam
	transitionToFire   bool
	currentStroke      ebiten_sprite.Stroke
}

func NewBattleState(level int) *BattleState {
	// loading assets, could be in init() and consistent usage of states.
	g := game.Start(data.HardcodedLevelGenerator{}, level)

	ebiten_canon.LoadBattleImages()

	ebg := ebiten_background.NewEbitenBattleGround(g.Battleground)
	emt := ebiten_monster.NewEbitenMonsterTeam(&g)

	bs := BattleState{
		game:               &g,
		ebitenBattleGround: ebg,
		ebitenMonsterTeam:  &emt,
		transitionToFire:   false,
		currentStroke:      nil,
	}

	ecd := ebiten_canon.NewEbitenCanonDeck(&g, bs.deployCannon, bs.moveCannon)
	bs.ebitenCanonDeck = &ecd

	return &bs
}

func (s *BattleState) Debug() string {
	return "battleState State"
}

func (s *BattleState) deployCannon(on int) {
	s.game.DeployCannon(on)
	s.transitionToFire = true
}

func (s *BattleState) moveCannon(from, to int) {
	s.game.MoveCannon(from, to)
	s.transitionToFire = true
}

func (s *BattleState) Update(stack *states.StateStack, keys []ebiten.Key) error {
	st, strokeStarted := ebiten_sprite.EbitenStrokeStarted()
	if strokeStarted {
		s.currentStroke = st
		s.ebitenCanonDeck.StrokeStart(st)
	}

	if s.currentStroke != nil {
		if s.currentStroke.IsJustReleased() {
			s.ebitenCanonDeck.ReleaseDrag()
		}
	}

	s.ebitenMonsterTeam.Update()
	s.ebitenCanonDeck.Update()

	if s.transitionToFire {
		s.transitionToFire = false
		fireState := NewFireCannonState(s.ebitenCanonDeck, s.ebitenMonsterTeam, s.game)
		stack.Push(fireState)
	}

	return nil
}

func (s *BattleState) Draw(screen *ebiten.Image) {
	s.ebitenBattleGround.Draw(screen)
	DrawGameOverLine(screen)
	s.ebitenCanonDeck.Draw(screen)
	s.ebitenMonsterTeam.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func DrawGameOverLine(screen *ebiten.Image) {
	lineColor := color.RGBA{R: 255, A: 255} // Red color

	vector.StrokeLine(screen, 0, GameOverLine,
		constants.ScreenWidth, GameOverLine,
		2, lineColor, false)
}
