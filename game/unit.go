package game

import (
	. "herzog/game/game_static"
	"herzog/lib/geometry"
)

type Unit struct {
	Code             int
	Faction          *Faction
	Action           UnitAction
	Order            Order
	CenterX, CenterY float64
	ChassisDegree    int
	Health, Fuel     int
	Turrets          []*Turret
}

func (u *Unit) GetPhysicalCenterCoords() (float64, float64) {
	return u.CenterX, u.CenterY
}

func (u *Unit) GetFaction() *Faction {
	return u.Faction
}

func (u *Unit) IsAlive() bool {
	return u.Health > 0
}

func (u *Unit) isInAir() bool {
	return false
}

func (u *Unit) GetTileCoords() (int, int) {
	return int(u.CenterX), int(u.CenterY)
}

func (u *Unit) GetStaticData() *UnitStatic {
	return STableUnits[u.Code]
}

func (u *Unit) rotateChassisTowardsVector(vx, vy float64) {
	degs := geometry.GetDegreeOfFloatVector(vx, vy)
	u.rotateTowardsDegree(degs)
}

func (u *Unit) snapTurretsDegreesToChassis() {
	for i := range u.Turrets {
		u.Turrets[i].RotationDegree = u.ChassisDegree
	}
}

func (u *Unit) rotateTowardsDegree(deg int) {
	if u.ChassisDegree == deg {
		return
	}
	rotateSpeed := geometry.GetDiffForRotationStep(u.ChassisDegree, deg, u.GetStaticData().ChassisRotationSpeed)
	u.ChassisDegree += rotateSpeed
	for i := range u.Turrets {
		u.Turrets[i].RotationDegree += rotateSpeed
	}
	u.normalizeDegrees()
}

func (u *Unit) normalizeDegrees() {
	u.ChassisDegree = geometry.NormalizeDegree(u.ChassisDegree)
	for i := range u.Turrets {
		u.Turrets[i].normalizeDegrees()
	}
}
