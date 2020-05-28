package main

import (
	"fmt"
	"github.com/evassilyev/demogame/core"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	units := InitUnits(core.TeamNum, core.UnitsInTeam)

	fmt.Print("Initial positions")
	core.ShowPositions(units)

	battle := core.NewBattle(units)
	battle.Start()

	fmt.Println("Final positions")
	core.ShowPositions(units)
}

func InitUnits(teamNumber, unitsInTeam int) (units []core.Unit) {
	teams := make([][]core.Unit, teamNumber)
	for i := 0; i < teamNumber; i++ {
		for j := 0; j < unitsInTeam; j++ {
			teams[i] = append(teams[i], core.NewUnit(rand.Intn(core.FieldSize), core.Team(i)))
		}
	}
	for i := 0; i < teamNumber; i++ {
		for _, u := range teams[i] {
			for j := 0; j < teamNumber; j++ {
				if j == i {
					continue
				}
				u.AddEnemies(teams[j])
			}
		}
	}
	for i := 0; i < teamNumber; i++ {
		units = append(units, teams[i]...)
	}
	return
}
