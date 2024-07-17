package game

import (
	. "herzog/game/game_static"
)

type Faction struct {
	Commander   *Commander
	HQBuilding  *Building
	ColorNumber int
	Gold        int

	currentBuiltUnitCode      int
	currentBuiltUnitOrderCode int
	currentBuildProgress      int
	buildsUnitNow             bool
}

func (f *Faction) GetTotalCostForCurrentProduction() int {
	unitCost := STableUnits[f.currentBuiltUnitCode].Cost
	orderCost := STableUnits[f.currentBuiltUnitCode].OrderCosts[f.currentBuiltUnitOrderCode]
	return unitCost + orderCost
}

func (f *Faction) ProductionInProgress() bool {
	return f.buildsUnitNow
}

func (f *Faction) FinishedProduction() bool {
	return f.currentBuildProgress == GetUnitStaticDataByCode(f.currentBuiltUnitCode).BuildTime
}

func (f *Faction) StartProduction() {
	f.buildsUnitNow = true
}

func (f *Faction) SetSelectedProduction(unitCode, orderCode int) {
	if f.buildsUnitNow {
		panic("Can't call this now")
	}
	f.currentBuiltUnitCode = unitCode
	f.currentBuiltUnitOrderCode = orderCode
}

func (f *Faction) GetSelectedProduction() (unitCode, orderCode int) {
	unitCode = f.currentBuiltUnitCode
	orderCode = f.currentBuiltUnitOrderCode
	return
}

func (f *Faction) GetCurrentProductionPercentage() int {
	return 100 * f.currentBuildProgress / GetUnitStaticDataByCode(f.currentBuiltUnitCode).BuildTime
}

func (f *Faction) ClearProductionState() {
	f.buildsUnitNow = false
	f.currentBuildProgress = 0
}

func (f *Faction) DoProductionStep() {
	cost := f.GetTotalCostForCurrentProduction()
	if f.FinishedProduction() {
		return
	}
	price := cost / GetUnitStaticDataByCode(f.currentBuiltUnitCode).BuildTime
	if f.currentBuildProgress == 0 { // Align price division on first spending tick
		price += cost % GetUnitStaticDataByCode(f.currentBuiltUnitCode).BuildTime
	}
	if f.Gold > price {
		f.Gold -= price
		f.currentBuildProgress++
	}
}
