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
