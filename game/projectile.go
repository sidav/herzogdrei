package game

import (
	. "herzog/game/game_static"
)

type Projectile struct {
	faction          *Faction
	staticData       *ProjectileStatic
	CenterX, CenterY float64
	RotationDegree   int
	fuel             float64 // how many 'speeds' it spends until it is destroyed
	whoShot          Actor
	targetActor      Actor // for homing projectiles
	SetToRemove      bool
	isInAir          bool
}

func (p *Projectile) GetStaticData() *ProjectileStatic {
	return p.staticData
}

func (p *Projectile) isHoming() bool {
	return p.GetStaticData().RotationSpeed > 0
}
