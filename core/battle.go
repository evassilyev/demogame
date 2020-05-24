package core

const FieldSize = 250

type Battle interface {
	Start()
}

type battle struct {
	units []Unit
}

func (b *battle) Start() {
	end := 30
	for {
		for _, u := range b.units {
			u.Aim()
		}
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
