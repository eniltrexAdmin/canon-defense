package pkg

import "github.com/hajimehoshi/ebiten/v2"

// State is coupled to Ebiten. TODO decouple it, if possible
type State interface {
	Update(stack *StateStack, keys []ebiten.Key) error
	Draw(screen *ebiten.Image)
	// TODO we can add init/enter/and exit or others.
}

type StateStack struct {
	states []State
}

func NewStateStack(initState State) StateStack {
	return StateStack{states: []State{initState}}
}

func (s *StateStack) Update(stack *StateStack, keys []ebiten.Key) error {
	lastState := s.states[len(s.states)-1]
	return lastState.Update(stack, keys)
}

func (s *StateStack) Draw(screen *ebiten.Image) {
	for _, state := range s.states {
		state.Draw(screen)
	}
}

func (s *StateStack) Push(state State) {
	s.states = append(s.states, state)
}

func (s *StateStack) Pop() {
	// there's always at least one state.
	if len(s.states) == 1 {
		return
	}
	s.states = s.states[:len(s.states)-1]
}

func (s *StateStack) Switch(state State) {
	s.states[len(s.states)-1] = state
}

func (s *StateStack) Clear() {
	s.states = []State{}
}
