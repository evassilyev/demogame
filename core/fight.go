package core

import (
	"fmt"
	"github.com/google/uuid"
)

type FightsHandler struct {
	infight map[string]string
	fights  map[string]map[string]Unit
}

func NewFightsHandler() *FightsHandler {
	fm := make(map[string]string)
	f := map[string]map[string]Unit{}
	return &FightsHandler{
		infight: fm,
		fights:  f,
	}
}

func (fh *FightsHandler) Involve(fighter Unit, enemy Unit) {
	fkey, fif := fh.infight[fighter.ID()]
	ekey, eif := fh.infight[enemy.ID()]
	if eif {
		fh.infight[fighter.ID()] = ekey
		if _, ok := fh.fights[ekey]; !ok {
			fh.fights[ekey] = map[string]Unit{}
		}
		fh.fights[ekey][fighter.ID()] = fighter
	} else {
		if !fif {
			fid := uuid.New().String()
			fh.infight[enemy.ID()] = fid
			fh.infight[fighter.ID()] = fid
			if _, ok := fh.fights[fid]; !ok {
				fh.fights[fid] = map[string]Unit{}
			}

			fh.fights[fid][enemy.ID()] = enemy
			fh.fights[fid][fighter.ID()] = fighter
		} else {
			fh.infight[enemy.ID()] = fkey
			if _, ok := fh.fights[fkey]; !ok {
				fh.fights[fkey] = map[string]Unit{}
			}
			fh.fights[fkey][enemy.ID()] = enemy
		}
	}
}

func (fh *FightsHandler) PrintFights() {
	fmt.Println("Current fights:")
	for key, val := range fh.fights {
		fmt.Print(key, ":")
		for _, unit := range val {
			if unit.IsAlive() {
				fmt.Print(unit.Team(), ":", unit.Health(), ",")
			}
		}
		fmt.Println()
	}
}

func (fh *FightsHandler) HandleFights() {
	for _, units := range fh.fights {
		for _, u := range units {
			if u.IsAlive() {
				for _, e := range units {
					if u.Team() != e.Team() && e.IsAlive() {
						e.GetDamage(u.Attack())
					}
				}
			}
		}
		for _, u := range units {
			u.Wound()
			if !u.IsAlive() {
				delete(fh.infight, u.ID())
			}
		}
	}

	var finised []string
	for key, units := range fh.fights {
		var u []Unit
		for _, v := range units {
			u = append(u, v)
		}

		f, _ := isFihised(u)
		if f {
			finised = append(finised, key)
		}
	}
	for _, key := range finised {
		delete(fh.fights, key)
	}
}

func isFihised(units []Unit) (res bool, winners Team) {
	winners = Team(-1)
	res = true
	first := true
	for _, u := range units {
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
