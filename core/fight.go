package core

import (
	"fmt"
	"github.com/google/uuid"
)

type FightsHandler struct {
	infight map[string]string
	fights  map[string][]Unit
}

func NewFightsHandler() *FightsHandler {
	fm := map[string]string{}
	f := map[string][]Unit{}
	return &FightsHandler{
		infight: fm,
		fights:  f,
	}
}

func (fh *FightsHandler) Involve(fighter Unit, enemy Unit) {
	fkey, fif := fh.infight[fighter.ID()]
	ekey, eif := fh.infight[enemy.ID()]
	if eif {
		if !fif {
			fh.infight[fighter.ID()] = ekey
			fh.fights[ekey] = append(fh.fights[ekey], fighter)
		}
	} else {
		if !fif {
			fid := uuid.New().String()
			fh.infight[enemy.ID()] = fid
			fh.infight[fighter.ID()] = fid
			fh.fights[fid] = append(fh.fights[fid], enemy, fighter)
		} else {
			fh.infight[enemy.ID()] = fkey
			fh.fights[fkey] = append(fh.fights[fkey], enemy)
		}
	}
}

func (fh *FightsHandler) HandleFights() {
	// TODO handle fights
}

func (fh *FightsHandler) PrintFights() {
	for key, val := range fh.fights {
		fmt.Print(key, ":")
		for _, unit := range val {
			fmt.Print(unit.Team())
		}
		fmt.Println()
	}
}
