package game

import (
	. "herzog/game/game_static"
	"herzog/lib/geometry"
)

type Commander struct {
	AsUnit               Unit
	CarriedUnit          *Unit
	IsFiring             bool
	IsTransforming       bool
	TransformingProgress int
}

func (c *Commander) GetPhysicalCenterCoords() (float64, float64) {
	return c.AsUnit.GetPhysicalCenterCoords()
}

func (c *Commander) GetFaction() *Faction {
	return c.AsUnit.Faction
}

func (c *Commander) IsAlive() bool {
	return c.AsUnit.Health > 0
}

func (c *Commander) isInAir() bool {
	return c.GetStaticData().IsAircraft
}

func (c *Commander) resetTransformation() {
	c.IsTransforming = false
	c.TransformingProgress = 0
}

func (c *Commander) GetTileCoordinates() (int, int) {
	return geometry.TrueCoordsToTileCoords(c.GetPhysicalCenterCoords())
}

func (c *Commander) GetStaticData() *UnitStatic {
	return STableUnits[c.AsUnit.Code]
}

// Action setters

func (c *Commander) ResetState() {
	c.AsUnit.Action.Reset()
}

func (c *Commander) SetMoveState(vx, vy float64) {
	c.AsUnit.Action.Kind = ACTION_CMOVE
	c.AsUnit.Action.Vx = vx
	c.AsUnit.Action.Vy = vy
}

func (c *Commander) SetPickupState() {
	c.AsUnit.Action.Kind = ACTION_CPICKUP
}

func (c *Commander) SetDropState() {
	c.AsUnit.Action.Kind = ACTION_CDROP
}
