package ai

import (
	"fmt"
	"herzog/game"
)

type AiStruct struct {
	btf     *game.Battlefield
	com     *game.Commander
	hq      *game.Building
	faction *game.Faction
	state   int

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
	a.decideProduction()
	if a.faction.Commander.IsAlive() {
		a.selectState()
		a.pilotCommander()
	}
}

func (a *AiStruct) cheatMoney() {
	if a.btf.CurrentTick%60 == 0 {
		a.faction.Gold += 10
	}
}

func (a *AiStruct) debugPrint() {
	fmt.Printf("=== TICK %d ===\n", a.btf.CurrentTick)
	if a.faction.ProductionInProgress() {
		fmt.Printf("$%d; producing %d%%... ", a.faction.Gold, a.faction.GetCurrentProductionPercentage())
	} else {
		fmt.Printf("$%d; standing by... ", a.faction.Gold)
	}
	fmt.Printf("Selected state: %d\n", a.state)
}
