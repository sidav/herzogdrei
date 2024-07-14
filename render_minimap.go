package main

import (
	"herzog/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *renderer) updateMinimap(b *game.Battlefield) {
	if r.minimap == nil {
		r.minimap = make([][]rl.Color, len(b.Tiles))
		for i := range b.Tiles {
			r.minimap[i] = make([]rl.Color, len(b.Tiles[0]))
		}
	}
	for x := range b.Tiles {
		for y := range b.Tiles[x] {
			actor := b.GetGroundActorAtTileCoordinates(x, y)
			if actor == nil {
				if b.Tiles[x][y].GetStaticData().CanBeWalkedOn {
					r.minimap[x][y] = rl.Black
				} else {
					r.minimap[x][y] = rl.Brown
				}
			} else {
				faction := actor.GetFaction()
				if faction == nil {
					r.minimap[x][y] = rl.LightGray
				} else {
					r.minimap[x][y] = *r.getFactionColor(faction)
				}
			}
		}
	}
}

func (r *renderer) renderMinimap(b *game.Battlefield, forFaction *game.Faction, atX, atY int32) {
	const OUTLINE = 2
	width := MINIMAP_PIXEL_SIZE * int32(len(r.minimap))
	height := MINIMAP_PIXEL_SIZE * int32(len(r.minimap[0]))
	r.drawOutlinedRect(atX-OUTLINE, atY-OUTLINE,
		width+OUTLINE*2,
		height+OUTLINE*2,
		2, rl.White, rl.Black)
	for x := range r.minimap {
		for y := range r.minimap[x] {
			if r.minimap[x][y] != rl.Black {
				rl.DrawRectangle(
					atX+(int32(x)*MINIMAP_PIXEL_SIZE),
					atY+(int32(y)*MINIMAP_PIXEL_SIZE),
					MINIMAP_PIXEL_SIZE, MINIMAP_PIXEL_SIZE,
					r.minimap[x][y],
				)
				// rl.DrawPixel(atX+int32(x), atY+int32(y), r.minimap[x][y])
			}
		}
	}
	cx, cy := forFaction.Commander.GetPhysicalCenterCoords()
	rl.DrawLine(atX+int32(cx*MINIMAP_PIXEL_SIZE), atY, atX+int32(cx*MINIMAP_PIXEL_SIZE), atY+height, *r.getFactionColor(forFaction))
	rl.DrawLine(atX, atY+int32(cy*MINIMAP_PIXEL_SIZE), atX+width, atY+int32(cy*MINIMAP_PIXEL_SIZE), *r.getFactionColor(forFaction))
}
