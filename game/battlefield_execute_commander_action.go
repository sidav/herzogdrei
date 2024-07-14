package game

import (
	. "herzog/game/game_static"
	"herzog/lib/geometry"
)

func (b *Battlefield) TryRefuelAndRepairCommander(com *Commander) {
	refuel := b.CurrentTick%1 == 0
	repair := b.CurrentTick%6 == 0
	if !(refuel || repair) {
		return
	}
	x, y := com.GetPhysicalCenterCoords()
	bld := b.GetBuildingAtTileCoordinates(int(x), int(y))
	if bld != nil && bld.Faction == com.AsUnit.Faction {
		if refuel {
			com.AsUnit.Fuel += 2
		}
		if repair {
			com.AsUnit.Health += 1
		}
	}
	if com.AsUnit.Fuel > com.GetStaticData().MaxFuel {
		com.AsUnit.Fuel = com.GetStaticData().MaxFuel
	}
	if com.AsUnit.Health > com.GetStaticData().MaxHitpoints {
		com.AsUnit.Health = com.GetStaticData().MaxHitpoints
	}
}

func (b *Battlefield) DoCommanderRespawnSequence(com *Commander) {
	com.CarriedUnit = nil
	const returnSpeed = 0.2
	hqx, hqy := com.AsUnit.Faction.HQBuilding.GetPhysicalCenterCoords()
	cx, cy := com.GetPhysicalCenterCoords()
	tx, ty := com.GetTileCoordinates()
	if b.GetGroundActorAtTileCoordinates(tx, ty) == com {
		b.Tiles[tx][ty].landActorHere = nil
	}
	if geometry.GetApproxDistFloat64(hqx, hqy, cx, cy) < returnSpeed {
		b.AddNewEffect(EFFECT_BIGGER_EXPLOSION, cx, cy, 7)
		com.AsUnit.Health = 1
		com.AsUnit.Fuel = com.GetStaticData().MaxFuel / 5
		com.AsUnit.ChassisDegree = 270
		com.AsUnit.snapTurretsDegreesToChassis()
		// Change to plane if we are on ground
		if !com.isInAir() {
			b.transformCommander(com)
		}
	} else {
		vx, vy := geometry.VectorToUnitVectorFloat64(hqx-cx, hqy-cy)
		com.AsUnit.CenterX += vx * returnSpeed
		com.AsUnit.CenterY += vy * returnSpeed
		if b.CurrentTick%15 == 0 {
			b.AddNewEffect(EFFECT_BIGGER_EXPLOSION, cx, cy, 0.75)
		}
	}
}

func (b *Battlefield) DoCommanderTransformationSequence(com *Commander) {
	tx, ty := com.GetTileCoordinates()
	if !(GetUnitStaticDataByCode(com.GetStaticData().TransformsTo).IsAircraft || b.AreCoordsPassable(tx, ty)) {
		com.resetTransformation()
		return
	}
	if com.AsUnit.ChassisDegree != 270 {
		com.AsUnit.rotateChassisTowardsVector(0, -1)
		return
	}
	if com.TransformingProgress == TICKS_FOR_TRANSFORMATION {
		b.transformCommander(com)
		return
	}
	com.TransformingProgress++
}

func (b *Battlefield) TryFireForCommander(com *Commander) {
	if !com.IsFiring {
		return
	}
	b.shootAsTurret(com, com.AsUnit.Turrets[0])
}

func (b *Battlefield) ExecuteCommanderAction(c *Commander) {
	if c.AsUnit.Fuel <= 0 {
		c.AsUnit.Health = 0
	}
	if !c.IsAlive() {
		b.DoCommanderRespawnSequence(c)
		return
	}
	if c.IsTransforming {
		b.DoCommanderTransformationSequence(c)
		return
	}
	b.TryRefuelAndRepairCommander(c)
	b.TryFireForCommander(c)
	switch c.AsUnit.Action.Kind {
	case ACTION_NONE:
		return
	case ACTION_CMOVE:
		b.ExecuteCMoveActionForCommander(c)
	case ACTION_CPICKUP:
		b.ExecuteCPickupActionForCommander(c)
	case ACTION_CDROP:
		b.ExecuteCDropActionForCommander(c)
	}
}

func (b *Battlefield) ExecuteCMoveActionForCommander(c *Commander) {
	vx, vy := geometry.VectorToUnitVectorFloat64(c.AsUnit.Action.Vx, c.AsUnit.Action.Vy)
	moveSpeed := c.GetStaticData().MovementSpeed

	targetDegree := geometry.GetDegreeOfFloatVector(vx, vy)
	if c.AsUnit.ChassisDegree != targetDegree {
		c.AsUnit.rotateChassisTowardsVector(vx, vy)
	} else {
		oldTx, oldTy := geometry.TrueCoordsToTileCoords(c.AsUnit.CenterX, c.AsUnit.CenterY)
		newCx, newCy := c.AsUnit.CenterX+vx*moveSpeed, c.AsUnit.CenterY+vy*moveSpeed
		newTx, newTy := geometry.TrueCoordsToTileCoords(newCx, newCy)
		if c.isInAir() || b.GetGroundActorAtTileCoordinates(newTx, newTy) == c ||
			b.AreCoordsPassable(newTx, newTy) {

			if !c.isInAir() {
				b.SwitchTilePointersForGroundActor(c, oldTx, oldTy, newTx, newTy)
			}
			c.AsUnit.CenterX = newCx
			c.AsUnit.CenterY = newCy
			spendFuel := 1
			if c.CarriedUnit != nil {
				spendFuel = 2
			}
			c.AsUnit.Fuel -= spendFuel
			if c.AsUnit.Fuel < 0 {
				c.AsUnit.Fuel = 0
			}
		}
	}
}

func (b *Battlefield) ExecuteCPickupActionForCommander(c *Commander) {
	tx, ty := c.GetTileCoordinates()
	bld := b.GetBuildingAtTileCoordinates(tx, ty)
	if bld != nil {
		if bld.Faction == c.AsUnit.Faction && bld.Faction.FinishedProduction() {
			// create the unit
			c.CarriedUnit = b.CreateNewUnit(
				c.AsUnit.Faction.CurrentBuiltUnitCode,
				c.AsUnit.Faction,
				c.AsUnit.CenterX,
				c.AsUnit.CenterY,
			)
			c.CarriedUnit.Order.Code = c.AsUnit.Faction.Commander.AsUnit.Faction.CurrentBuiltUnitOrderCode
			// Reset the faction build state
			c.AsUnit.Faction.ClearProductionState()
		}
	} else {
		u := b.TakeGroundUnitFromTileCoordinates(tx, ty)
		if u != nil {
			c.CarriedUnit = u
		}
	}
	c.AsUnit.Action.Kind = ACTION_NONE
}

func (b *Battlefield) ExecuteCDropActionForCommander(c *Commander) {
	tx, ty := c.GetTileCoordinates()
	if b.AreCoordsPassable(tx, ty) {
		c.CarriedUnit.ChassisDegree = c.AsUnit.ChassisDegree
		c.CarriedUnit.snapTurretsDegreesToChassis()
		c.CarriedUnit.CenterX, c.CarriedUnit.CenterY = geometry.TileCoordsToTrueCoords(tx, ty)

		c.CarriedUnit.Action.Kind = ACTION_NONE
		c.CarriedUnit.Order.SetOrigin(tx, ty)

		b.addActor(c.CarriedUnit)
		c.CarriedUnit = nil

		c.AsUnit.Action.Kind = ACTION_NONE
	}
}

func (b *Battlefield) transformCommander(com *Commander) {
	tx, ty := com.GetTileCoordinates()
	com.AsUnit.Code = com.GetStaticData().TransformsTo
	com.AsUnit.Turrets[0].staticData = com.GetStaticData().TurretsData[0]
	com.AsUnit.snapTurretsDegreesToChassis()
	com.resetTransformation()
	actorAtTile := b.GetGroundActorAtTileCoordinates(tx, ty)
	// If transformed TO plane...
	if com.isInAir() {
		if actorAtTile == com {
			b.Tiles[tx][ty].landActorHere = nil
		} else if actorAtTile == com.GetFaction().HQBuilding {
			// It's OK, do nothing
		} else if actorAtTile != com {
			// Something is wrong
			panic("Transformation collision error")
		}
	} else {
		if actorAtTile != nil {
			panic("Transformation collision error")
		}
		b.Tiles[tx][ty].landActorHere = com
	}
}
