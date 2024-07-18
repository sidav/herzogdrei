package ai

import "herzog/game/game_static"

func (a *AiStruct) decideProduction() {
	if !a.faction.ProductionInProgress() {
		a.faction.SetSelectedProduction(
			game_static.UNIT_INFANTRY,
			game_static.ORDER_CAPTURE,
		)
		a.faction.StartProduction()
	}
}
