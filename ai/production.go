package ai

import "herzog/game/game_static"

func (a *AiStruct) decideProduction() {
	if a.faction.ProductionInProgress() {
		return
	}
	unitCode := game_static.UNIT_INFANTRY
	unitOrder := game_static.ORDER_CAPTURE
	if a.targetBuilding.GetFaction() != nil || rnd.Rand(10) == 0 {
		unitCode, unitOrder = a.selectRandomProduction()
	}
	a.faction.SetSelectedProduction(unitCode, unitOrder)
	a.faction.StartProduction()
}

func (a *AiStruct) selectRandomProduction() (int, int) {
	code := 0
	for game_static.STableUnits[code].IsCommander {
		code = rnd.Rand(len(game_static.STableUnits))
	}
	order := -1
	for order == -1 || !game_static.STableUnits[code].CanDoOrder(order) {
		order = rnd.Rand(len(game_static.STableUnits[code].OrderCosts))
	}
	return code, order
}
