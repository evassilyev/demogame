package core

const FieldSize = 250

type Battle interface {
	Start()
}

type battle struct {
	units []Unit
}

func (b *battle) Start() {
	end := 30 // TODO remove test condition

	for {
		for _, u := range b.units {
			target := u.NearestEnemy()
			var d = target.Position() - u.Position()
			if d <= 1 && d >= -1 {
				// TODO load unit to fight
				continue
			}
			u.Aim(d)
		}

		// TODO Handle all fights

		for _, u := range b.units {
			u.Move()
		}

		end--
		if 0 > end {
			break
		}
	}
}

func NewBattle(u []Unit) Battle {
	return &battle{
		units: u,
	}
}
