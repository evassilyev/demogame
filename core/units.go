package core

import (
	"github.com/google/uuid"
)

type Unit interface {
	Position() int
	Team() Team
	Enemies() []Unit
	AddEnemies([]Unit)
	Aim()
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

func (u *unit) nearestEnemy() (nearest Unit) {
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

func (u *unit) Aim() {
	target := u.nearestEnemy()
	u.moveToTarget(target)
}

func (u *unit) Enemies() []Unit {
	return u.enemies
}

func (u *unit) AddEnemies(e []Unit) {
	u.enemies = append(u.enemies, e...)
}

func (u *unit) moveToTarget(u2 Unit) {
	var d = u2.Position() - u.position
	if d == 1 || d == -1 || d == 0 {
		// do nothing (fight)
		return
	}
	if d > 0 {
		u.newPosition = u.position + 1
	} else {
		u.newPosition = u.position - 1
	}
}

func (u *unit) Position() int {
	return u.position
}

func (u *unit) Team() Team {
	return u.team
}

func NewUnit(position int, team Team) Unit {
	var u unit
	u.id = uuid.New().String()
	u.position = position
	u.newPosition = position
	u.team = team
	return &u
}
