package main

import (
	"herzog/game"
	"herzog/lib/geometry"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const MINIMAP_PIXEL_SIZE = 3

var DEFAULT_TINT = rl.RayWhite
var FOG_OF_WAR_TINT = rl.Color{
	R: 96,
	G: 96,
	B: 96,
	A: 255,
}

type renderer struct {
	camTopLeftX, camTopLeftY int32
	viewportW, viewportH     int32
	minimapRenderTextureMask rl.Texture2D
	btl                      *game.Battlefield
	minimap                  [][]rl.Color

	// for technical/debug output
	lastFrameRenderingTime time.Duration
}

func (r *renderer) renderBattlefield(b *game.Battlefield) {
	r.btl = b
	timeFrameRenderStarted := time.Now()
	r.viewportW = WINDOW_W
	r.viewportH = WINDOW_H
	r.updateCameraCoordinates(b.Factions[0].Commander.GetPhysicalCenterCoords())

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	for x := range b.Tiles {
		for y := range b.Tiles[x] {
			r.renderTile(b, x, y)
		}
	}

	// render ground units
	for _, a := range b.Actors {
		if u, ok := a.(*game.Unit); ok {
			r.renderUnit(b, u)
		}
	}
	// render buildings
	for _, a := range b.Actors {
		if bld, ok := a.(*game.Building); ok {
			r.renderBuilding(b, nil, bld)
		}
	}
	// render commander
	for _, a := range b.Actors {
		if com, ok := a.(*game.Commander); ok {
			if a.IsAlive() {
				r.renderCommander(b, com)
			}
		}
	}
	// render projectiles
	for _, p := range b.Projectiles {
		r.renderProjectile(p)
	}

	// // render aircrafts
	// for i := b.units.Front(); i != nil; i = i.Next() {
	// 	if i.Value.(*unit).getStaticData().IsAircraft {
	// 		r.renderUnit(b, pc, i.Value.(*unit))
	// 	}
	// }

	for _, e := range b.Effects {
		r.renderEffect(b, e)
	}

	r.lastFrameRenderingTime = time.Since(timeFrameRenderStarted) / time.Millisecond

	r.renderUI(b, b.Factions[0])
	rl.EndDrawing()
}

func (r *renderer) renderTile(b *game.Battlefield, x, y int) {
	t := b.Tiles[x][y]
	osx, osy := r.physicalToOnScreenCoords(float64(x), float64(y))
	if !r.isRectInViewport(osx, osy, TILE_SIZE_IN_PIXELS, TILE_SIZE_IN_PIXELS) {
		return
	}
	tintToUse := FOG_OF_WAR_TINT
	if t.GetStaticData().BackgroundSpriteCode != "" {
		rl.DrawTexture(
			tilesAtlaces[game.STableTiles[t.Code].BackgroundSpriteCode].GetSpriteByFrame(0),
			osx,
			osy,
			tintToUse,
		)
	}
	rl.DrawTexture(
		tilesAtlaces[game.STableTiles[t.Code].SpriteCodes[t.SpriteVariantIndex]].GetSpriteByFrame(0),
		osx,
		osy,
		tintToUse,
	)
}

func (r *renderer) renderFactionFlagAt(btl *game.Battlefield, f *game.Faction, leftX, bottomY int32) {
	frame := (6 * btl.CurrentTick / 25) % uiAtlaces["factionflag"].TotalFrames()
	spr := uiAtlaces["factionflag"].GetSpriteByColorAndFrame(f.ColorNumber, frame)
	rl.DrawTexture(
		spr,
		leftX,
		bottomY-spr.Height,
		DEFAULT_TINT,
	)
}

func (r *renderer) updateCameraCoordinates(pcx, pcy float64) {
	r.camTopLeftX, r.camTopLeftY = int32(pcx*TILE_SIZE_IN_PIXELS), int32(pcy*TILE_SIZE_IN_PIXELS)
	r.camTopLeftX -= WINDOW_W / 2
	r.camTopLeftY -= WINDOW_H / 2
}

func (r *renderer) physicalToOnScreenCoords(physX, physY float64) (int32, int32) {
	pixx, pixy := r.physicalToPixelCoords(physX, physY)
	pixx -= r.camTopLeftX
	pixy -= r.camTopLeftY
	return int32(pixx), int32(pixy)
}

func (r *renderer) isRectInViewport(rx, ry, rw, rh int32) bool {
	return geometry.AreTwoCellRectsOverlapping32(rx, ry, rw, rh, 0, 0, r.viewportW, r.viewportH)
}

func (r *renderer) AreOnScreenCoordsInViewport(osx, osy int32) bool {
	// log.Printf("%d, %d \n", osx, osy)
	return osx >= 0 && osx < int32(r.viewportW) && osy >= 0 && osy < int32(r.viewportH)
}

func (r *renderer) physicalToPixelCoords(px, py float64) (int32, int32) {
	return int32(float32(px) * PIXEL_TO_PHYSICAL_RATIO), int32(float32(py) * PIXEL_TO_PHYSICAL_RATIO)
}
