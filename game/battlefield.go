package game

import (
	"herzog/lib/astar"
	"herzog/lib/geometry"
)

type Battlefield struct {
	Tiles       [][]tile
	Factions    []*Faction
	Actors      []Actor
	Projectiles []*Projectile
	Effects     []*Effect
	CurrentTick int

	pathfinder *astar.AStarPathfinder
}

func (b *Battlefield) AreTileCoordsValid(tx, ty int) bool {
	return tx >= 0 && tx < len(b.Tiles) && ty >= 0 && ty < len(b.Tiles[0])
}

func (b *Battlefield) areActorsInRangeFromEachOther(a1, a2 Actor, rang float64) bool {
	a1x, a1y := a1.GetPhysicalCenterCoords()
	a2x, a2y := a2.GetPhysicalCenterCoords()
	return geometry.GetApproxDistFloat64(a1x, a1y, a2x, a2y) <= rang
}

func (b *Battlefield) getApproxRangeBetweenCoordinatables(a1, a2 Coordinatable) float64 {
	a1x, a1y := a1.GetPhysicalCenterCoords()
	a2x, a2y := a2.GetPhysicalCenterCoords()
	return geometry.GetApproxDistFloat64(a1x, a1y, a2x, a2y)
}

func (b *Battlefield) AreCoordsPassable(tx, ty int) bool {
	return b.AreTileCoordsValid(tx, ty) &&
		b.Tiles[tx][ty].GetStaticData().CanBeWalkedOn && b.GetGroundActorAtTileCoordinates(tx, ty) == nil
}

func (b *Battlefield) GetGroundActorAtTileCoordinates(tx, ty int) Actor {
	if b.AreTileCoordsValid(tx, ty) {
		return b.Tiles[tx][ty].landActorHere
	}
	return nil
}

func (b *Battlefield) GetAirActorAtRealCoordinates(x, y float64) Actor {
	for _, a := range b.Actors {
		ax, ay := a.GetPhysicalCenterCoords()
		if a.isInAir() && geometry.GetApproxDistFloat64(x, y, ax, ay) < 0.5 {
			return a
		}
	}
	return nil
}

func (b *Battlefield) GetBuildingAtTileCoordinates(tx, ty int) *Building {
	for _, a := range b.Actors {
		if bld, ok := a.(*Building); ok {
			if bld.IsPresentAt(tx, ty) {
				return bld
			}
		}
	}
	return nil
}

func (b *Battlefield) GetGroundUnitAtTileCoordinates(tx, ty int) *Unit {
	if !b.AreTileCoordsValid(tx, ty) {
		return nil
	}
	a := b.Tiles[tx][ty].landActorHere
	if u, ok := a.(*Unit); ok {
		return u
	}
	return nil
}

// Both returns the unit AND removes it from the battlefield. Useful for transport picking up
func (b *Battlefield) TakeGroundUnitFromTileCoordinates(tx, ty int) *Unit {
	if !b.AreTileCoordsValid(tx, ty) {
		return nil
	}
	a := b.Tiles[tx][ty].landActorHere
	if u, ok := a.(*Unit); ok {
		b.Tiles[tx][ty].landActorHere = nil
		b.removeActorFromList(a)
		return u
	}
	return nil
}

func (b *Battlefield) addActor(a Actor) {
	if a == nil {
		panic("Nil actor!")
	}
	if u, ok := a.(*Unit); ok {
		if !u.GetStaticData().IsAircraft {
			tx, ty := geometry.TrueCoordsToTileCoords(u.GetPhysicalCenterCoords())
			b.Tiles[tx][ty].landActorHere = a
		}
	}
	if bld, ok := a.(*Building); ok {
		for tx := bld.TopLeftX; tx < bld.TopLeftX+bld.GetStaticData().W; tx++ {
			for ty := bld.TopLeftY; ty < bld.TopLeftY+bld.GetStaticData().H; ty++ {
				b.Tiles[tx][ty].landActorHere = a
			}
		}
	}
	b.Actors = append(b.Actors, a)
}

func (b *Battlefield) removeActorFromList(act Actor) {
	for i, a := range b.Actors {
		if a == act {
			b.Actors = append(b.Actors[:i], b.Actors[i+1:]...)
			return
		}
	}
}
