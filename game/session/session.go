package session

type Session struct {
	CompletedLevels []bool
}

func NewSession(totalLevels int) Session {
	return Session{
		CompletedLevels: make([]bool, totalLevels),
	}
}

func (s *Session) CompleteLevel(level int) {
	if level >= 0 && level < len(s.CompletedLevels) {
		s.CompletedLevels[level] = true
	}
}
