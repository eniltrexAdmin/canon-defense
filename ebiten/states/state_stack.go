package states

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

// State is coupled to Ebiten. TODO decouple it, if possible
type State interface {
	Update(stack *StateStack, keys []ebiten.Key) error
	Draw(screen *ebiten.Image)
	Debug() string
	//Init()// TODO we can add init/enter/and exit or others.
}

type StateStack struct {
	states []State
}

func NewStateStack() StateStack {
	return StateStack{states: []State{}}
}

func (s *StateStack) Update(stack *StateStack, keys []ebiten.Key) error {
	lastState := s.states[len(s.states)-1]
	return lastState.Update(stack, keys)
}

func (s *StateStack) Draw(screen *ebiten.Image) {
	for _, state := range s.states {
		//fmt.Printf("Drawing: %s\n", state.Debug())
		state.Draw(screen)
	}
}

func (s *StateStack) Push(state State) {
	s.states = append(s.states, state)
	s.Debug()
}

func (s *StateStack) Pop() {
	// there's always at least one state.
	if len(s.states) == 1 {
		return
	}
	s.states = s.states[:len(s.states)-1]
	s.Debug()
}

func (s *StateStack) Switch(state State) {
	// This function should take care of init(). It should call the init() function in a go subroutine
	// and draw the loading screen until it's not finished. (or create a loading state that does all the
	// img, _, err := image.Decode(bytes.NewReader(images.Tiles_png))
	//

	s.states[len(s.states)-1] = state
	s.Debug()
}

func (s *StateStack) Clear() {
	s.states = []State{}
}

func (s *StateStack) Debug() {
	println("Debug:---")
	for position, state := range s.states {
		fmt.Printf("Pos: %d: %s\n", position, state.Debug())
	}
	println("====")
}
