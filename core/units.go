package core

import (
	"github.com/google/uuid"
)

type Unit interface {
	ID() string
	Position() int
	Team() Team
	AddEnemies([]Unit)
	NearestEnemy() Unit
	Aim(int)
	Move()
}

type Team int

type unit struct {
	id          string
	position    int
	newPosition int
	team        Team
	enemies     []Unit
}

func (u *unit) Move() {
	u.position = u.newPosition
}

func (u *unit) NearestEnemy() (nearest Unit) {
	distance := FieldSize + 1
	for _, e := range u.enemies {
		d := e.Position() - u.position
		if d < 0 {
			d = -d
		}
		if d < distance {
			distance = d
			nearest = e
		}
	}
	return
}

func (u *unit) Aim(d int) {
	if d > 0 {
		u.newPosition = u.position + 1
	} else {
		u.newPosition = u.position - 1
	}
}

func (u *unit) AddEnemies(e []Unit) {
	u.enemies = append(u.enemies, e...)
}

func (u *unit) Position() int {
	return u.position
}

func (u *unit) Team() Team {
	return u.team
}

func (u *unit) ID() string {
	return u.id
}

func NewUnit(position int, team Team) Unit {
	var u unit
	u.id = uuid.New().String()
	u.position = position
	u.newPosition = position
	u.team = team
	return &u
}
