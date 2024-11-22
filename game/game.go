package game

import "fmt"

type CanonBuilder struct{}

func (cd CanonBuilder) Create() Canon {
	return BuildCanon(1)
}

type Turn int

type CanonTDGame struct {
	Battleground Battleground
	CanonDeck    CanonDeck
	MonsterTeam  MonsterTeam
	CanonBuilder CanonBuilder
	Turns        []UserActions
	playerTurn   bool
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

func (g *CanonTDGame) CurrentTurn() Turn {
	println(fmt.Printf("Turns: %+v\n", g.Turns))
	return Turn(len(g.Turns))
}

func (g *CanonTDGame) DeployCannon(column int) {
	if !g.playerTurn {
		panic("its not player turn!")
	}
	c := g.CanonBuilder.Create()
	de := g.CanonDeck.deployCannon(BattleGroundColumn(column), &c)
	g.Turns = append(g.Turns, de)
	g.playerTurn = false
}

func (g *CanonTDGame) MoveCannon(origin, destination int) {
	if !g.playerTurn {
		panic("its not player turn!")
	}
	de, err := g.CanonDeck.MoveCanon(BattleGroundColumn(origin), BattleGroundColumn(destination))
	if err != nil {
		panic(err)
	}
	g.Turns = append(g.Turns, de)
	g.playerTurn = false
}

func (g *CanonTDGame) HitMonster(c *Canon, m *Monster) {
	println(fmt.Sprintf("about to damage %d to monster with life %d on turn %d", c.Damage, m.HealthPoints, g.CurrentTurn()))
	m.Hit(c, g.CurrentTurn())
	g.printBattleGround()
}

func (g *CanonTDGame) MonstersCharge() {
	if g.playerTurn {
		panic("its not monster turn!")
	}
	g.MonsterTeam.charge()
	g.playerTurn = true
}

func (g *CanonTDGame) printBattleGround() {
	for _, monster := range g.MonsterTeam.Monsters {
		println(fmt.Sprintf("one Monster with max life %d and current life: %d", monster.MaxLife, monster.HealthPoints))
	}
}

// Probably this one is just "Server" not needed
//func (g *CanonTDGame) Fire() {
//	for column, c := range g.CanonDeck.Canons {
//		if c != nil {
//			g.MonsterTeam.DamageMonsters(c, column)
//		}
//	}
//}

type BattleStarted struct {
	DomainEvent
}
