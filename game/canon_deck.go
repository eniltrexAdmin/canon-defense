package game

import (
	"errors"
	"fmt"
	"time"
)

type CanonDeck struct {
	Canons        map[BattleGroundColumn]*Canon
	canonCapacity int
}

func NewCanonDeck(b Battleground) CanonDeck {
	canons := make(map[BattleGroundColumn]*Canon, b.Columns)
	return CanonDeck{
		Canons:        canons,
		canonCapacity: int(b.Columns),
	}
}

func (cd *CanonDeck) CanonCapacity() int {
	return cd.canonCapacity
}

func (cd *CanonDeck) deployCannon(position BattleGroundColumn, canon *Canon) CanonDeployed {
	if cd.Canons[position] != nil {
		cd.Canons[position].merge(canon)
	} else {
		cd.Canons[position] = canon
	}
	return CanonDeployed{
		DomainEvent: DomainEvent{occurredOn: time.Now().Format("2006-01-02 15:04:05")},
		Canon:       *canon,
		Destination: position,
	}
}

type CanonDeployed struct {
	DomainEvent
	Canon       Canon
	Destination BattleGroundColumn
}

func (cd *CanonDeck) MoveCanon(origin, destination BattleGroundColumn) (CanonMoved, error) {
	println(fmt.Sprintf("Moving cannon from %d to %d", origin, destination))
	if origin == destination {
		return CanonMoved{}, nil
	}
	canon := cd.Canons[origin]
	if canon == nil {
		return CanonMoved{}, errors.New("no canon to move here")
	}

	if cd.Canons[destination] == nil {
		cd.Canons[destination] = canon
	} else {
		cd.Canons[destination].merge(canon)
	}
	delete(cd.Canons, origin)
	return CanonMoved{
		DomainEvent: DomainEvent{occurredOn: time.Now().Format("2006-01-02 15:04:05")},
		Canon:       *canon,
		Origin:      origin,
		Destination: destination,
	}, nil
}

type CanonMoved struct {
	DomainEvent
	Canon       Canon
	Origin      BattleGroundColumn
	Destination BattleGroundColumn
}
