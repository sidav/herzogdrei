package ai

import (
	"herzog/lib/random"
)

var rnd random.PRNG

func SetPRNG(r random.PRNG) {
	rnd = r
}

func (a *AiStruct) getCommanderTileCoords() (int, int) {
	return a.faction.Commander.GetTileCoordinates()
}

func (a *AiStruct) getCommanderRealCoords() (float64, float64) {
	return a.faction.Commander.GetPhysicalCenterCoords()
}

func (a *AiStruct) isHealthOrFuelCritical() bool {
	return a.com.AsUnit.GetFuelPercentage() < 20 || a.com.AsUnit.GetHpPercentage() < 33
}

func (a *AiStruct) areHealthAndFuelFull() bool {
	return a.com.AsUnit.GetFuelPercentage() == 100 && a.com.AsUnit.GetHpPercentage() == 100
}

func (a *AiStruct) commanderArrived(tx, ty int) bool {
	x, y := a.faction.Commander.GetTileCoordinates()
	return tx == x && ty == y
}

func floatVectorForTileCoords(fx, fy, tx, ty int) (float64, float64) {
	return float64(tx - fx), float64(ty - fy)
}
