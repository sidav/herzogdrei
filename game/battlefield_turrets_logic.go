package game

import "herzog/lib/geometry"

func (b *Battlefield) ActForAllActorsTurrets(a Actor) {
	if u, ok := a.(*Unit); ok {
		for i := range u.Turrets {
			if u.Turrets[i].nextTickToAct > b.CurrentTick {
				continue
			}
			b.actTurret(a, u.Turrets[i])
		}
	}
}

func (b *Battlefield) actTurret(shooter Actor, t *Turret) {
	// shooterTileX, shooterTileY := 0, 0
	shooterX, shooterY := 0.0, 0.0
	turretRange := t.GetStaticData().FireRange
	var orderedTarget Actor
	if u, ok := shooter.(*Unit); ok {
		shooterX, shooterY = u.GetPhysicalCenterCoords()
		orderedTarget = u.Order.TargetActor
	} else {
		panic("Unknown shooter type!")
	}
	if orderedTarget != nil && orderedTarget.GetFaction() != shooter.GetFaction() {
		t.targetActor = orderedTarget
	}
	// Cancel the current turret's target if it can't be attacked right now
	if t.targetActor != nil {
		if !t.targetActor.IsAlive() ||
			!b.areActorsInRangeFromEachOther(shooter, t.targetActor, turretRange) ||
			!b.isTargetAttackableForTurret(t, t.targetActor) {

			t.targetActor = nil
		}
	}

	// if targetActor not set...
	if t.targetActor == nil {
		t.targetActor = b.getGoodTargetForActorsTurret(shooter, t, false, false)
	}
	// if still no target, return
	if t.targetActor == nil {
		t.nextTickToAct = b.CurrentTick + 10 // sleep for some time
		return
	}
	var targetCenterX, targetCenterY float64
	rotateTo := 0
	targetCenterX, targetCenterY = t.targetActor.GetPhysicalCenterCoords()
	rotateTo = geometry.GetDegreeOfFloatVector(targetCenterX-shooterX, targetCenterY-shooterY)

	if t.canRotate() {
		t.RotationDegree += geometry.GetDiffForRotationStep(t.RotationDegree, rotateTo, t.GetStaticData().RotateSpeed)
		t.normalizeDegrees()
	}

	if abs(t.RotationDegree-rotateTo) <= t.GetStaticData().FireSpreadDegrees/2 {
		b.shootAsTurret(shooter, t)
	}
}

func (b *Battlefield) shootAsTurret(shooter Actor, t *Turret) {
	if t.nextTickToAct > b.CurrentTick {
		return
	}
	shooterX, shooterY := shooter.GetPhysicalCenterCoords()

	vectX, vectY := geometry.DegreeToUnitVector(t.RotationDegree)
	projX, projY := shooterX+vectX/2, shooterY+vectY/2 // TODO: turret displacement
	degreeSpread := rnd.RandInRange(-t.GetStaticData().FireSpreadDegrees, t.GetStaticData().FireSpreadDegrees)
	rangeSpread := t.GetStaticData().ShotRangeSpread * float64(rnd.RandInRange(-100, 100)) / 100

	proj := &Projectile{
		faction:        shooter.GetFaction(),
		staticData:     t.GetStaticData().FiredProjectileData,
		CenterX:        projX,
		CenterY:        projY,
		RotationDegree: t.RotationDegree + degreeSpread,
		whoShot:        shooter,
		targetActor:    t.targetActor,
	}

	var projectileFuel float64
	var isInAir bool
	if t.targetActor == nil {
		projectileFuel = t.GetStaticData().FireRange + rangeSpread
		isInAir = shooter.isInAir()
	} else {
		targetCenterX, targetCenterY := t.targetActor.GetPhysicalCenterCoords()
		if proj.isHoming() {
			projectileFuel = 1.5 * (t.GetStaticData().FireRange + rangeSpread)
		} else {
			projectileFuel = geometry.GetPreciseDistFloat64(targetCenterX, targetCenterY, shooterX, shooterY) + rangeSpread - 0.5
		}
		isInAir = t.targetActor.isInAir()
	}
	proj.fuel = projectileFuel
	proj.isInAir = isInAir

	b.Projectiles = append(b.Projectiles, proj)

	if t.GetStaticData().MaxShotsInVolley > 1 {
		t.shotsInCurrentVolley++
		if t.shotsInCurrentVolley < t.GetStaticData().MaxShotsInVolley {
			t.nextTickToAct = b.CurrentTick + t.GetStaticData().CooldownPerShot
			return
		} else {
			t.shotsInCurrentVolley = 0
		}
	}
	t.nextTickToAct = b.CurrentTick + t.GetStaticData().CooldownAfterVolley
}

func (b *Battlefield) isTargetAttackableForTurret(t *Turret, target Actor) bool {
	if target.isInAir() {
		return t.staticData.AttacksAir
	}
	return t.staticData.AttacksLand
}

func (b *Battlefield) getGoodTargetForActorsTurret(a Actor, t *Turret, ignoreRange, allowBuildings bool) Actor {
	var currTarget Actor
	for _, enemy := range b.Actors {
		bld, isBuilding := enemy.(*Building)
		if isBuilding && !(allowBuildings && bld.IsAttackable()) {
			continue
		}
		if !enemy.IsAlive() || enemy.GetFaction() == a.GetFaction() || (enemy.isInAir() && !t.staticData.AttacksAir) ||
			(!enemy.isInAir() && !t.staticData.AttacksLand) {
			continue
		}
		if ignoreRange || b.areActorsInRangeFromEachOther(a, enemy, t.GetStaticData().FireRange) {
			if currTarget == nil ||
				b.getApproxRangeBetweenCoordinatables(a, currTarget) >= b.getApproxRangeBetweenCoordinatables(a, enemy) {
				currTarget = enemy
			}
		}
	}
	return currTarget
}
