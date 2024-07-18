package game

import (
	. "herzog/game/game_static"
	"herzog/lib/geometry"
)

func (b *Battlefield) ExecuteUnitOrder(u *Unit) {
	if u.Order.TargetActor != nil && !u.Order.TargetActor.IsAlive() {
		u.Order.TargetActor = nil
	}
	switch u.Order.Code {
	case ORDER_STANDBY:
		b.executeStandbyOrderForUnit(u)
	case ORDER_PATROL:
		b.executePatrolOrderForUnit(u)
	case ORDER_CAPTURE:
		b.executeCaptureOrderForUnit(u)
	case ORDER_SEARCHNDESTROY:
		b.executeSeachAndDestroyOrderForUnit(u)
	default:
		panic("Unimplemented order")
	}
}

func (b *Battlefield) executeStandbyOrderForUnit(u *Unit) {
	if len(u.Turrets) == 0 || u.Action.Kind != ACTION_NONE {
		return
	}
	if u.Order.TargetActor == nil || !b.areActorsInRangeFromEachOther(u, u.Order.TargetActor, u.Turrets[0].GetStaticData().FireRange) {
		u.Order.TargetActor = b.getGoodTargetForActorsTurret(u, u.Turrets[0], false, true)
	}
	if u.Order.TargetActor != nil {
		u.Action.Kind = ACTION_ROTATE
		u.Action.Vx, u.Action.Vy = u.Order.TargetActor.GetPhysicalCenterCoords()
		u.Action.Vx -= u.CenterX
		u.Action.Vy -= u.CenterY
	}
}

func (b *Battlefield) executePatrolOrderForUnit(u *Unit) {
	const patrolRadius = 5
	for u.Order.TargetX == 0 && u.Order.TargetY == 0 ||
		!b.AreCoordsPassable(u.Order.GetTargetTileCoords()) {

		u.Order.TargetX, u.Order.TargetY = rnd.RandInRange(u.Order.OriginX-patrolRadius, u.Order.OriginX+patrolRadius),
			rnd.RandInRange(u.Order.OriginY-patrolRadius, u.Order.OriginY+patrolRadius)
	}
	b.selectTargetAndMoveToItForUnitOrder(u, false, false)
	if u.Order.TargetActor != nil {

	} else {
		b.setUnitStateForPath(u, u.Order.TargetX, u.Order.TargetY)
	}
}

func (b *Battlefield) executeCaptureOrderForUnit(u *Unit) {
	tx, ty := u.GetTileCoords()
	if u.Order.TargetActor == nil {
		ucx, ucy := u.GetPhysicalCenterCoords()
		// Find closest uncaptured building
		distance := 65535.0 // arbitrary big integer
		var targetBuilding Actor
		for _, a := range b.Actors {
			if bld, ok := a.(*Building); ok {
				if bld.Faction != u.Faction {
					bldcx, bldcy := bld.GetPhysicalCenterCoords()
					bldDistance := geometry.GetApproxDistFloat64(ucx, ucy, bldcx, bldcy)
					if bldDistance < distance {
						distance = bldDistance
						targetBuilding = a
					}
				}
			}
		}
		u.Order.TargetActor = targetBuilding
	}
	if u.Order.TargetActor != nil {
		if u.Order.TargetActor.GetFaction() == u.GetFaction() {
			u.Order.TargetActor = nil
			return
		}
		bld := u.Order.TargetActor.(*Building)
		bx, by := bld.TopLeftX, bld.TopLeftY
		if geometry.GetApproxDistFromTo(tx, ty, bx, by) <= 1 { // Enter building
			u.Action.Kind = ACTION_ENTER_BUILDING
			u.Action.SetVectorByInt(bx-tx, by-ty)
			u.Action.CalcMoveRemaining()
		} else {
			b.setUnitStateForPath(u, bx, by)
		}
	}
}

func (b *Battlefield) executeSeachAndDestroyOrderForUnit(u *Unit) {
	if len(u.Turrets) == 0 || u.Action.Kind != ACTION_NONE {
		panic("SnD order for unit without turrets!")
	}
	b.selectTargetAndMoveToItForUnitOrder(u, true, false)
}

func (b *Battlefield) selectTargetAndMoveToItForUnitOrder(u *Unit, ignoreRange, allowBuildings bool) {
	if u.Order.TargetActor == nil ||
		!u.Order.TargetActor.IsAlive() ||
		!b.areActorsInRangeFromEachOther(u, u.Order.TargetActor, u.Turrets[0].GetStaticData().FireRange) {

		u.Order.TargetActor = b.getGoodTargetForActorsTurret(u, u.Turrets[0], ignoreRange, allowBuildings)
	}
	if u.Order.TargetActor != nil {
		requiredRange := u.Turrets[0].staticData.FireRange
		if u.Turrets[0].canRotate() {
			requiredRange = 2 * u.Turrets[0].staticData.FireRange / 3
		}
		if !b.areActorsInRangeFromEachOther(u, u.Order.TargetActor, requiredRange) {
			tx, ty := geometry.TrueCoordsToTileCoords(u.Order.TargetActor.GetPhysicalCenterCoords())
			b.setUnitStateForPath(u, tx, ty)
		} else {
			u.Action.Kind = ACTION_ROTATE
			u.Action.Vx, u.Action.Vy = u.Order.TargetActor.GetPhysicalCenterCoords()
			u.Action.Vx -= u.CenterX
			u.Action.Vy -= u.CenterY
		}
	}
}

func (b *Battlefield) setUnitStateForPath(u *Unit, targetX, targetY int) {
	utx, uty := u.GetTileCoords()
	vx, vy := b.getVectorForPath(utx, uty, targetX, targetY)
	u.Action.Kind = ACTION_START_MOVING
	u.Action.SetVectorByInt(vx, vy)
	u.Action.CalcMoveRemaining()
}
