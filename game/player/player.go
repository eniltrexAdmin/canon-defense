package player

type Player struct {
	CurrentLevel int
}

func NewPlayer() Player {
	return Player{
		CurrentLevel: 2,
	}
}
