package game

import (
	. "herzog/game/game_static"
	"herzog/lib/geometry"
)

const unitsForCapture = 4 // This many units are required to capture a building
type Building struct {
	Faction            *Faction
	code               int
	Hitpoints          int
	TopLeftX, TopLeftY int                       // tile coords
	CaptureProgress    [unitsForCapture]*Faction // array of faction pointers
}

func (b *Building) GetStaticData() *BuildingStatic {
	return STableBuildings[b.code]
}

func (b *Building) GetFaction() *Faction {
	return b.Faction
}

func (b *Building) GetTileCenterCoordinates() (int, int) {
	return b.TopLeftX + b.GetStaticData().W/2, b.TopLeftY + b.GetStaticData().H/2
}

func (b *Building) IsAlive() bool {
	return true
}

func (b *Building) IsAttackable() bool {
	return b.GetStaticData().MaxHitpoints > 0
}

func (b *Building) isInAir() bool {
	return false
}

func (b *Building) GetPhysicalCenterCoords() (float64, float64) {
	return float64(b.TopLeftX) + float64(b.GetStaticData().W)/2, float64(b.TopLeftY) + float64(b.GetStaticData().H)/2
}

func (b *Building) IsPresentAt(tileX, tileY int) bool {
	w, h := b.GetStaticData().W, b.GetStaticData().H
	return geometry.AreCoordsInTileRect(tileX, tileY, b.TopLeftX, b.TopLeftY, w, h)
}

func (b *Building) IsFullyCaptured() bool {
	var f *Faction = b.CaptureProgress[0]
	for i := range b.CaptureProgress {
		if b.CaptureProgress[i] != f {
			return false
		}
	}
	return true
}
