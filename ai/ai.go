package ai

import (
	"herzog/game"
)

type AiStruct struct {
	btf     *game.Battlefield
	com     *game.Commander
	hq      *game.Building
	faction *game.Faction
	state   int

	// macro-related variables
	targetBuilding *game.Building

	// state-related variables
	coordsSet bool
	tx, ty    int
}

func CreateNewAiStruct(b *game.Battlefield, faction *game.Faction) *AiStruct {
	return &AiStruct{
		btf:     b,
		faction: faction,
		com:     faction.Commander,
		hq:      faction.HQBuilding,
	}
}

func (a *AiStruct) Act() {
	a.cheatMoney()
	a.decideMacro()
	a.decideProduction()
	a.selectState()
	if a.faction.Commander.IsAlive() {
		a.pilotCommander()
	}
}

func (a *AiStruct) cheatMoney() {
	if a.btf.CurrentTick%60 == 0 {
		a.faction.Gold += 10
	}
}

func (a *AiStruct) debugPrint() {
	printf("=== TICK %d ===\n", a.btf.CurrentTick)
	if a.faction.ProductionInProgress() {
		printf("$%d; producing %d%%... ", a.faction.Gold, a.faction.GetCurrentProductionPercentage())
	} else {
		printf("$%d; standing by... ", a.faction.Gold)
	}
	printf("Selected state: %d\n", a.state)
	tx, ty := a.targetBuilding.GetPhysicalCenterCoords()
	printf("Target building at %.1f, %.1f\n", tx, ty)
}
