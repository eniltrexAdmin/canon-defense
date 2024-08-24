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

func (g CanonTDGame) PlaceCannon(position int) {
	column := ToBattleGroundColumn(position, g.Battleground)
	c := BuildCanon(canonDamage(1))
	err := g.CanonDeck.placeCanon(column, &c)
	if err != nil {
		panic(err.Error())
	}
	// plus do FIRE!!
}
