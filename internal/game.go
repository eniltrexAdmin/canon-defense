package internal

func main() {
	lg := LevelGenerator{}
	bg := lg.Generate()
	cd := NewCanonDeck(bg)

	// everything ready to start the battle,
	for {
		canon := NewCanon()
		err := cd.placeCanon(3, canon)
		if err != nil {
			return
		} // or move cannon
		bg.fire(cd)
		bg.move()
		// check if game over
	}
}
