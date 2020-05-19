package core

type unit struct {
	health   int
	position int
	weapon   Weapon
	enemies  []Unit
}

type Unit interface {
	MoveToTarget(Unit)
	Position() int
	Attack(Unit)
}

type Weapon interface {
	GetDamage() int
}

type weapon struct {
	damage func() int
}
