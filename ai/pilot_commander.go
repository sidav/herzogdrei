package ai

import "herzog/game"

func (a *AiStruct) pilotCommander() {
	switch a.state {
	case aiStateUndecided, aiStateRefuelRepair:
		a.refuelRepairCommander()
	case aiStatePickUpBuiltUnit:
		a.pickUpBuiltUnit()
	case aiStateDropBuiltUnit:
		a.dropBuiltUnit()
	}
}

func (a *AiStruct) refuelRepairCommander() {
	if a.areHealthAndFuelFull() {
		a.resetState()
	} else {
		x, y := a.getClosestRefuelCoords() // TODO: use coords setting
		a.flyCommanderToCoords(x, y)
	}
}

func (a *AiStruct) pickUpBuiltUnit() {
	if a.com.CarriedUnit != nil {
		a.resetState()
		a.state = aiStateDropBuiltUnit
		return
	}
	x, y := a.getClosestBuiltUnitPickupCoords() // TODO: use coords setting
	if !a.commanderArrived(x, y) {
		a.flyCommanderToCoords(x, y)
		return
	}
	a.com.SetPickupState()
}

func (a *AiStruct) dropBuiltUnit() {
	if !a.coordsSet || !a.btf.AreCoordsPassable(a.tx, a.ty) {
		var dropNear game.Actor = a.targetBuilding
		dropRange := 5
		if a.isHealthOrFuelCritical() {
			dropNear = a.com // If fuel is critical, drop ASAP!
			dropRange = 4
		} else if a.targetBuilding.GetFaction() != nil {
			// if target bld is not neutral, drop near com
			//  TODO: drop near closest base?
			dropNear = a.com
			dropRange = 7
		}
		a.setCoords(a.getUnitDropCoords(dropNear, dropRange))
	}
	if !a.commanderArrived(a.tx, a.ty) {
		a.flyCommanderToCoords(a.tx, a.ty)
		return
	}
	a.com.SetDropState()
	a.resetState()
}

// Helper func
func (a *AiStruct) flyCommanderToCoords(x, y int) {
	if a.commanderArrived(x, y) {
		a.com.ResetState()
		return
	}
	cx, cy := a.getCommanderTileCoords()

	a.com.SetMoveState(floatVectorForTileCoords(
		cx, cy, x, y,
	))
}
