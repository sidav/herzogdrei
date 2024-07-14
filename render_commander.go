package main

import (
	. "herzog/game"
	"herzog/game/game_static"
	"herzog/lib/geometry"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *renderer) renderCommander(b *Battlefield, c *Commander) {
	if c.IsTransforming && c.TransformingProgress > 0 {
		completenessPrecent := (100*c.TransformingProgress + game_static.TICKS_FOR_TRANSFORMATION/2) / game_static.TICKS_FOR_TRANSFORMATION
		x, y := c.GetPhysicalCenterCoords()
		osx, osy := r.physicalToOnScreenCoords(x, y)
		if !r.isRectInViewport(osx, osy, TILE_SIZE_IN_PIXELS, TILE_SIZE_IN_PIXELS) {
			return
		}
		neededAtlas := effectsAtlaces[c.GetStaticData().TransformationAnimationSpritesCode]
		currentFrame := geometry.GetPartitionIndex(completenessPrecent, 0, 100, neededAtlas.TotalFrames())
		rl.DrawTexture(
			neededAtlas.GetSpriteByFrame(currentFrame),
			osx-neededAtlas.GetSpriteByFrame(currentFrame).Width/2,
			osy-neededAtlas.GetSpriteByFrame(currentFrame).Height/2,
			DEFAULT_TINT, // proj.faction.factionColor,
		)
	} else if c.IsAlive() {
		r.renderUnit(b, &c.AsUnit)
	}
}
