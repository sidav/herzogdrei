package game

import (
	. "herzog/game/game_static"
	"herzog/lib/random"
)

const (
	STARTING_MONEY = 2500
)

func (b *Battlefield) Init(r random.PRNG) {
	SetPRNG(r)

	b.Factions = make([]*Faction, 0)
	b.initFromStringMap()
}

func (b *Battlefield) initAndPlaceNewFaction(hqx, hqy int) {
	newFact := &Faction{
		Gold:        STARTING_MONEY,
		ColorNumber: len(b.Factions),
	}
	b.Factions = append(b.Factions, newFact)
	b.CreateCommanderForFaction(newFact, float64(hqx)+1.5, float64(hqy)+1.5)
	bld := &Building{
		code:     BLD_MAIN_BASE,
		TopLeftX: hqx,
		TopLeftY: hqy,
		Faction:  newFact,
	}
	newFact.HQBuilding = bld
	bld.Hitpoints = bld.GetStaticData().MaxHitpoints
	b.addActor(bld)
}
