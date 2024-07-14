package game

import (
	. "herzog/game/game_static"
)

type Faction struct {
	Commander   *Commander
	HQBuilding  *Building
	ColorNumber int
	Gold        int

	CurrentBuiltUnitCode      int
	CurrentBuiltUnitOrderCode int
	CurrentBuildProgress      int
	BuildsUnitNow             bool
}

func (f *Faction) GetTotalCostForCurrentProduction() int {
	unitCost := STableUnits[f.CurrentBuiltUnitCode].Cost
	orderCost := STableUnits[f.CurrentBuiltUnitCode].OrderCosts[f.CurrentBuiltUnitOrderCode]
	return unitCost + orderCost
}

func (f *Faction) FinishedProduction() bool {
	return f.CurrentBuildProgress == GetUnitStaticDataByCode(f.CurrentBuiltUnitCode).BuildTime
}

func (f *Faction) GetCurrentProductionPercentage() int {
	return 100 * f.CurrentBuildProgress / GetUnitStaticDataByCode(f.CurrentBuiltUnitCode).BuildTime
}

func (f *Faction) ClearProductionState() {
	f.BuildsUnitNow = false
	f.CurrentBuildProgress = 0
}

func (f *Faction) DoProductionStep() {
	cost := f.GetTotalCostForCurrentProduction()
	if f.FinishedProduction() {
		return
	}
	price := cost / GetUnitStaticDataByCode(f.CurrentBuiltUnitCode).BuildTime
	if f.CurrentBuildProgress == 0 { // Align price division on first spending tick
		price += cost % GetUnitStaticDataByCode(f.CurrentBuiltUnitCode).BuildTime
	}
	if f.Gold > price {
		f.Gold -= price
		f.CurrentBuildProgress++
	}
}
