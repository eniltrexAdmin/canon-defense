package game

type CanonBuilder struct{}

func (cd CanonBuilder) Create() Canon {
	return BuildCanon(1)
}

type CanonTDGame struct {
	Battleground Battleground
	CanonDeck    CanonDeck
	MonsterTeam  MonsterTeam
	CanonBuilder CanonBuilder
	Turns        []UserActions
}

type LevelGenerator interface {
	Generate(level int) CanonTDGame
}

func Start(level int) CanonTDGame {
	return HardcodedLevelGenerator{}.Generate(level)
}

type UserActions interface {
	OccurredOn() string
}

type DomainEvent struct {
	occurredOn string
}

func (e DomainEvent) OccurredOn() string {
	return e.occurredOn
}

func (g *CanonTDGame) DeployCannon(column int) {
	c := g.CanonBuilder.Create()
	de := g.CanonDeck.deployCannon(BattleGroundColumn(column), &c)
	g.Turns = append(g.Turns, de)
}

func (g *CanonTDGame) MoveCannon(origin, destination int) {
	de, err := g.CanonDeck.MoveCanon(BattleGroundColumn(origin), BattleGroundColumn(destination))
	if err != nil {
		panic(err)
	}
	g.Turns = append(g.Turns, de)
}

func (g *CanonTDGame) finishTurn() {

}
