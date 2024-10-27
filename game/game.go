package game

type CanonTDGame struct {
	Battleground Battleground
	CanonDeck    CanonDeck
}

func Start(level int) CanonTDGame {

	bg := HardcodedLevelGenerator{}.Generate(level)
	canonDeck := NewCanonDeck(bg)

	return CanonTDGame{
		Battleground: bg,
		CanonDeck:    canonDeck,
	}
}
