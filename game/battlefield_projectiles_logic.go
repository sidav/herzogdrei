package game

import (
	"herzog/lib/geometry"
	"math"
)

func (b *Battlefield) ActForProjectile(p *Projectile) {
	if p.SetToRemove {
		return // workaround for emptying the list
	}
	// move forward
	vx, vy := geometry.DegreeToUnitVector(p.RotationDegree)
	spd := math.Min(p.GetStaticData().Speed, p.fuel)
	p.CenterX += spd * vx
	p.CenterY += spd * vy
	p.fuel -= spd

	var hitTarget Actor
	if p.targetActor != nil { // Shot by unit
		if p.isHoming() {
			targX, targY := p.targetActor.GetPhysicalCenterCoords()
			rotateTo := geometry.GetDegreeOfFloatVector(targX-p.CenterX, targY-p.CenterY)
			p.RotationDegree += geometry.GetDiffForRotationStep(p.RotationDegree, rotateTo, p.GetStaticData().RotationSpeed)
			p.RotationDegree = geometry.NormalizeDegree(p.RotationDegree)
			if geometry.GetApproxDistFloat64(targX, targY, p.CenterX, p.CenterY) < 0.5 {
				hitTarget = p.targetActor
				p.SetToRemove = true
			}
		}
		if p.fuel <= 0 && hitTarget == nil {
			if p.isInAir {
				hitTarget = b.GetAirActorAtRealCoordinates(p.CenterX, p.CenterY)
			} else {
				tilex, tiley := geometry.TrueCoordsToTileCoords(p.CenterX, p.CenterY)
				hitTarget = b.GetGroundActorAtTileCoordinates(tilex, tiley)
				if hitTarget != nil && !(hitTarget.IsAlive() && hitTarget.isInAir() == p.targetActor.isInAir()) {
					hitTarget = nil
				}
			}
			p.SetToRemove = true
		}
	} else { // Shot by player
		if p.isInAir {
			hitTarget = b.GetAirActorAtRealCoordinates(p.CenterX, p.CenterY)
		} else {
			tilex, tiley := geometry.TrueCoordsToTileCoords(p.CenterX, p.CenterY)
			hitTarget = b.GetGroundActorAtTileCoordinates(tilex, tiley)
			if hitTarget != nil {
				if !(hitTarget.IsAlive() && hitTarget.isInAir() == p.isInAir) ||
					hitTarget == p.whoShot ||
					b.getApproxRangeBetweenCoordinatables(p, hitTarget) > 0.5+p.GetStaticData().Size {

					hitTarget = nil
				}
			}
		}
		if p.fuel <= 0 || hitTarget != nil {
			p.SetToRemove = true
		}
	}
	if p.SetToRemove {
		if hitTarget != nil {
			b.dealDamage(p.staticData.HitDamage, hitTarget)
		}
		if p.GetStaticData().CreatesEffectOnImpact {
			b.Effects = append(b.Effects,
				&Effect{
					CenterX:            p.CenterX,
					CenterY:            p.CenterY,
					SplashCircleRadius: p.GetStaticData().SplashRadius,
					Code:               p.GetStaticData().EffectCreatedOnImpactCode,
					CreationTick:       b.CurrentTick,
				})
		}
	}
	// debugWritef("%+v spd: %f\n", p, spd)
}
