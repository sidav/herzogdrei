package game

type Actor interface {
	GetPhysicalCenterCoords() (float64, float64)
	GetFaction() *Faction
	IsAlive() bool
	isInAir() bool
}
