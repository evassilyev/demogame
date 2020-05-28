package core

import "fmt"

const FieldSize = 250

type Battle interface {
	Start()
	IsFinished() (bool, Team)
}

type battle struct {
	units []Unit
	fh    *FightsHandler
}

func (b *battle) Start() {

	i := 0

	for {
		i++
		if i%10 == 0 {
			ShowPositions(b.units)
		}

		for _, u := range b.units {
			if u.IsAlive() {
				target := u.NearestEnemy()
				var d = target.Position() - u.Position()
				if d <= 1 && d >= -1 {
					b.fh.Involve(u, target)
					continue
				}
				u.Aim(d)
			}
		}
		b.fh.HandleFights()

		finished, winners := b.IsFinished()
		if finished {
			b.PrintWinners(winners)
			break
		}

		for _, u := range b.units {
			if u.IsAlive() {
				u.Move()
			}
		}
	}
	b.fh.PrintFights()
}

func (b *battle) IsFinished() (res bool, winners Team) {
	winners = Team(-1)
	res = true
	first := true
	for _, u := range b.units {
		if u.IsAlive() {
			if first {
				winners = u.Team()
				first = false
			}
			if u.Team() != winners {
				res = false
				return
			}
		}
	}
	return
}

func NewBattle(u []Unit) Battle {
	f := NewFightsHandler()
	return &battle{
		fh:    f,
		units: u,
	}
}

func (b *battle) PrintWinners(winners Team) {
	if winners == Team(-1) {
		fmt.Println("All units are dead")
	} else {
		var wa int
		for _, u := range b.units {
			if u.IsAlive() && u.Team() == winners {
				wa++
			}
		}
		fmt.Println(fmt.Sprintf("%d units from team %d won", wa, int(winners)))
	}
}
