package core

import "fmt"

const (
	UnitsInTeam = 10
	TeamNum     = 3
)

func ShowPositions(units []Unit) {
	fmt.Println()
	for t := 0; t < TeamNum; t++ {
		for p := 0; p <= FieldSize; p++ {
			unit := false
			for _, u := range units {
				if u.Team() == Team(t) && u.Position() == p && u.IsAlive() {
					fmt.Print(t)
					unit = true
				}
			}
			if unit {
				continue
			}
			fmt.Print("_")
		}
		fmt.Println()
	}
}
