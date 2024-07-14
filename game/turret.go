package game

import (
	. "herzog/game/game_static"
	"herzog/lib/geometry"
)

type Turret struct {
	staticData           *TurretStatic
	RotationDegree       int
	nextTickToAct        int
	shotsInCurrentVolley int

	targetActor              Actor
	targetTileX, targetTileY int
}

func (t *Turret) canRotate() bool {
	return t.GetStaticData().RotateSpeed > 0
}

func (t *Turret) GetStaticData() *TurretStatic {
	return t.staticData
}

func (t *Turret) normalizeDegrees() {
	t.RotationDegree = geometry.NormalizeDegree(t.RotationDegree)
}
