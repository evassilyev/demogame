package main

import (
	"fmt"
	"github.com/evassilyev/demogame/core"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	units := initUnits(core.FieldSize, teams, unitsInTeam)
	showPositions(units)
	battle := core.NewBattle(units)
	battle.Start()
	fmt.Println("After move")
	showPositions(units)
}

const (
	unitsInTeam = 10
	teams       = 2
)

func initUnits(field, tn, uit int) (units []core.Unit) {
	teams := make([][]core.Unit, tn)
	for i := 0; i < tn; i++ {
		for j := 0; j < uit; j++ {
			teams[i] = append(teams[i], core.NewUnit(rand.Intn(field), core.Team(i)))
		}
	}
	for i := 0; i < tn; i++ {
		for _, u := range teams[i] {
			for j := 0; j < tn; j++ {
				if j == i {
					continue
				}
				u.AddEnemies(teams[j])
			}
		}
	}
	for i := 0; i < tn; i++ {
		units = append(units, teams[i]...)
	}
	return
}

func showPositions(units []core.Unit) {
	for t := 0; t < teams; t++ {
		for p := 0; p <= core.FieldSize; p++ {
			unit := false
			for _, u := range units {
				if u.Team() == core.Team(t) && u.Position() == p {
					fmt.Print(t)
					unit = true
				}
			}
			if unit {
				continue
			}
			fmt.Print("_")
		}
		fmt.Print("\n")
	}
}
