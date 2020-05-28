package core

const FieldSize = 250

type Battle interface {
	Start()
}

type battle struct {
	units []Unit
	fh    *FightsHandler
}

func (b *battle) Start() {
	end := 30 // TODO remove test condition

	for {
		for _, u := range b.units {
			target := u.NearestEnemy()
			var d = target.Position() - u.Position()
			if d <= 1 && d >= -1 {
				b.fh.Involve(u, target)
				continue
			}
			u.Aim(d)
		}
		// TODO Handle all fights
		// b.fh.HandleFights()
		// TODO break loop condition - units of just one team alive or all units dead

		for _, u := range b.units {
			u.Move()
		}

		end--
		if 0 > end {
			break
		}
	}
	b.fh.PrintFights()
}

func NewBattle(u []Unit) Battle {
	f := NewFightsHandler()
	return &battle{
		fh:    f,
		units: u,
	}
}
