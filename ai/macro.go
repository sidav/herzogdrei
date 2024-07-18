package ai

import (
	"herzog/game"
	"herzog/lib/geometry"
)

func (ai *AiStruct) decideMacro() {
	if ai.targetBuilding == nil || ai.targetBuilding.GetFaction() == ai.faction {
		ai.selectNewTargetBuilding()
	}
}

func (ai *AiStruct) selectNewTargetBuilding() {
	fromX, fromY := ai.hq.GetPhysicalCenterCoords()
	bld := ai.btf.SelectActorWithHighestScore(
		func(a game.Actor) (score int, selectable bool) {
			b, ok := a.(*game.Building)
			if !ok || a.GetFaction() == ai.faction {
				return 0, false
			}
			// Set the score up
			bx, by := b.GetPhysicalCenterCoords()
			// The higher, the better
			score = 0
			if b.GetFaction() == nil {
				score += 100
			}
			score -= int(geometry.GetApproxDistFloat64(fromX, fromY, bx, by))
			score += 10 * rnd.Rand(5)
			return score, true
		},
	)
	ai.targetBuilding = bld.(*game.Building)
}
