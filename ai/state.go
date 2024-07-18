package ai

const (
	aiStateUndecided = iota
	aiStateRefuelRepair
	aiStatePickUpBuiltUnit
	aiStateDropBuiltUnit
)

func (a *AiStruct) resetState() {
	a.state = aiStateUndecided
	a.coordsSet = false
}

func (a *AiStruct) setCoords(x, y int) {
	a.tx = x
	a.ty = y
	a.coordsSet = true
}

func (a *AiStruct) selectState() {
	if !a.com.IsAlive() {
		a.state = aiStateRefuelRepair
	}
	if a.state == aiStateDropBuiltUnit && a.com.CarriedUnit == nil {
		panic("Something is wrong")
	}
	if a.state != aiStateUndecided {
		return
	}
	if a.isHealthOrFuelCritical() {
		a.state = aiStateRefuelRepair
	} else if a.faction.IsProductionFinished() {
		a.state = aiStatePickUpBuiltUnit
	}
	a.debugPrint()
}
