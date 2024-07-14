package game

import (
	. "herzog/game/game_static"
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

func (c *Commander) transform() {
	c.AsUnit.Code = c.GetStaticData().TransformsTo
	c.AsUnit.Turrets[0].staticData = c.GetStaticData().TurretsData[0]
	c.AsUnit.snapTurretsDegreesToChassis()
	c.resetTransformation()
}

func (c *Commander) GetTileCoordinates() (int, int) {
	x, y := c.GetPhysicalCenterCoords()
	return int(x), int(y)
}

func (c *Commander) GetStaticData() *UnitStatic {
	return STableUnits[c.AsUnit.Code]
}
