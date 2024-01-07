package pkg

type State interface {
	Update() error
	Draw()
	// TODO we can add init/enter/and exit or others.
}

type StateStack struct {
	states []State
}

func NewStateStack(initState State) StateStack {
	return StateStack{states: []State{initState}}
}

func (s *StateStack) Update() error {
	lastState := s.states[len(s.states)-1]
	return lastState.Update()
}

func (s *StateStack) Draw() {
	for _, state := range s.states {
		state.Draw()
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

func (s *StateStack) Clear() {
	s.states = []State{}
}
