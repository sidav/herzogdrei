package game

import (
	. "herzog/game/game_static"
)

type Effect struct {
	CenterX, CenterY   float64
	Code               EffectCode
	CreationTick       int
	SplashCircleRadius float64
}

func (e *Effect) GetStaticData() *EffectStatic {
	return STableEffects[e.Code]
}

func (e *Effect) GetExpirationPercent(currentTick int) int {
	return 100 * (currentTick - e.CreationTick) / e.GetStaticData().DefaultLifeTime
}
