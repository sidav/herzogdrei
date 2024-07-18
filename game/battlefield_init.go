package game

import (
	. "herzog/game/game_static"
	"herzog/lib/random"
)

func (b *Battlefield) Init(r random.PRNG) {
	SetPRNG(r)

	b.Factions = make([]*Faction, 0)
	b.initFromStringMap()

	// Random enemy units for debug
	// f1bx, f1by := b.Factions[0].HQBuilding.TopLeftX, b.Factions[0].HQBuilding.TopLeftY
	// for i := 0; i < 20; i++ {
	// 	x, y := -1, -1
	// 	for !b.AreCoordsPassable(x, y) || geometry.GetApproxDistFromTo(x, y, f1bx, f1by) < 7 {
	// 		x, y = rnd.Rand(len(b.Tiles)), rnd.Rand(len(b.Tiles[0]))
	// 	}
	// 	rx, ry := geometry.TileCoordsToTrueCoords(x, y)
	// 	u := b.CreateNewUnit(
	// 		rnd.RandInRange(UNIT_QUAD, UNIT_TANK),
	// 		b.Factions[1],
	// 		rx, ry,
	// 	)
	// 	u.ChassisDegree = 45 * rnd.Rand(8)
	// 	u.snapTurretsDegreesToChassis()
	// 	b.addActor(u)
	// }
}

func (b *Battlefield) initAndPlaceNewFaction(hqx, hqy int) {
	newFact := &Faction{
		Gold:        1000,
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
