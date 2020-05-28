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
	fmt.Println("Battle conditions: ")
	fmt.Printf("Teams: %d\nUnits in Team: %d\n", TeamNum, UnitsInTeam)

	i := 0

	for {
		i++
		if i%5 == 0 {
			ShowPositions(b.units)
			b.fh.PrintFights()
			/*
				for _, u := range b.units {
					if u.IsAlive() {
						fmt.Println(u.Team(), ":", u.Health(), "-", u.NearestEnemy().Team(),":", u.NearestEnemy().Health())
					}
				}
			*/
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
}

func (b *battle) IsFinished() (res bool, winners Team) {
	return isFihised(b.units)
}

func NewBattle(u []Unit) Battle {
	f := NewFightsHandler()
	return &battle{
		fh:    f,
		units: u,
	}
}

func (b *battle) PrintWinners(winners Team) {
	fmt.Println()
	for i := 0; i < FieldSize; i++ {
		fmt.Print("*")
	}
	fmt.Println()

	if winners == Team(-1) {
		fmt.Println("All units are dead")
	} else {
		var wa int
		var wnrs []Unit
		for _, u := range b.units {
			if u.IsAlive() && u.Team() == winners {
				wa++
				wnrs = append(wnrs, u)
			}
		}
		fmt.Println(fmt.Sprintf("%d units from team %d won", wa, int(winners)))
		for _, u := range wnrs {
			fmt.Println("Unit: ", u.ID(), " ", "Health:", u.Health())
		}
		fmt.Println()
	}
}
