package player

type Player struct {
	CurrentLevel int
}

func NewPlayer() Player {
	return Player{
		CurrentLevel: 1,
	}
}

func (p *Player) LevelCompleted() {
	p.CurrentLevel++
}
