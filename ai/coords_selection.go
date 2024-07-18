package ai

import (
	"herzog/game"
	"herzog/lib/geometry"
)

func (a *AiStruct) getClosestRefuelCoords() (int, int) {
	var selected *game.Building
	cx, cy := a.getCommanderRealCoords()
	for _, b := range a.btf.Actors {
		if b.GetFaction() == a.faction {
			if bld, ok := b.(*game.Building); ok {
				if selected == nil {
					selected = bld
					continue
				}
				bx, by := b.GetPhysicalCenterCoords()
				selX, selY := selected.GetPhysicalCenterCoords()
				if geometry.GetApproxDistFloat64(cx, cy, selX, selY) > geometry.GetApproxDistFloat64(cx, cy, bx, by) {
					selected = bld
				}
			}
		}
	}
	if selected != nil {
		return geometry.TrueCoordsToTileCoords(selected.GetPhysicalCenterCoords())
	}
	panic("Bad logic")
}

func (ai *AiStruct) getClosestBuiltUnitPickupCoords() (int, int) {
	// Pick up from the building closest to the target one
	cx, cy := ai.targetBuilding.GetPhysicalCenterCoords()
	bld := ai.btf.SelectActorWithHighestScore(
		func(act game.Actor) (int, bool) {
			b, ok := act.(*game.Building)
			if !ok || act.GetFaction() != ai.faction {
				return 0, false
			}
			bx, by := b.GetPhysicalCenterCoords()
			return -int(geometry.GetApproxDistFloat64(cx, cy, bx, by)), true
		},
	)
	if bld != nil {
		return geometry.TrueCoordsToTileCoords(bld.GetPhysicalCenterCoords())
	}
	panic("Bad logic")
}

func (a *AiStruct) getUnitDropCoords(searchNear game.Actor, searchRange int) (int, int) {
	// TODO: better logic regarding to the unit's order
	tx, ty := geometry.TrueCoordsToTileCoords(searchNear.GetPhysicalCenterCoords())
	tries := 0
	for tries < 25 {
		tries++
		x, y := rnd.RandInRange(tx-searchRange, tx+searchRange), rnd.RandInRange(ty-searchRange, ty+searchRange)
		if a.btf.AreCoordsPassable(x, y) {
			return x, y
		}
	}
	panic("Bad logic")
}
