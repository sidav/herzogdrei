package main

import (
	. "herzog/game"
	"herzog/lib/geometry"
	spritesatlas "herzog/lib/sprites_atlas"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *renderer) renderCommander(b *Battlefield, c *Commander) {
	r.renderUnit(b, &c.AsUnit)
}

func (r *renderer) renderUnit(b *Battlefield, u *Unit) {
	x, y := u.GetPhysicalCenterCoords()
	offset := 0.5
	chassisW := unitChassisAtlaces["commander"].GetSpriteByColorDegreeAndFrameNumber(0, u.ChassisDegree, 0).Width
	if chassisW > TILE_SIZE_IN_PIXELS {
		offset = float64(chassisW) / float64(TILE_SIZE_IN_PIXELS) / 2
	}
	osx, osy := r.physicalToOnScreenCoords(x-offset, y-offset)
	// log.Printf("%d, %d \n", osx, osy)
	if !r.isRectInViewport(osx, osy, TILE_SIZE_IN_PIXELS, TILE_SIZE_IN_PIXELS) {
		return
	}

	// draw chassis sprite
	rl.DrawTexture(
		unitChassisAtlaces[u.GetStaticData().ChassisSpriteCode].
			GetSpriteByColorDegreeAndFrameNumber(u.Faction.ColorNumber, u.ChassisDegree, 0),
		osx,
		osy,
		DEFAULT_TINT,
	)

	// draw turrets
	for turrIndex := range u.Turrets {
		if u.Turrets[turrIndex].GetStaticData().SpriteCode == "" {
			continue
		}
		usd := u.GetStaticData()
		// calculate turret displacement
		dsplX, dsplY := usd.TurretsData[turrIndex].TurretCenterX, usd.TurretsData[turrIndex].TurretCenterY
		if dsplX != 0 || dsplY != 0 {
			// rotate according to units face
			chassisShownDegree := geometry.SnapDegreeToFixedDirections(u.ChassisDegree, 8)
			dsplX, dsplY = geometry.RotateFloat64Vector(dsplX, dsplY, chassisShownDegree)
		}
		turrOsX, turrOsY := r.physicalToOnScreenCoords(x+dsplX, y+dsplY)
		sprite := turretsAtlaces[u.Turrets[turrIndex].GetStaticData().SpriteCode].GetSpriteByColorDegreeAndFrameNumber(u.Faction.ColorNumber, u.Turrets[turrIndex].RotationDegree, 0)
		rl.DrawTexture(
			sprite,
			turrOsX-sprite.Width/2,
			turrOsY-sprite.Height/2,
			DEFAULT_TINT,
		)
	}
	if u.Health < u.GetStaticData().MaxHitpoints {
		r.drawProgressBar(osx, osy+TILE_SIZE_IN_PIXELS-2, TILE_SIZE_IN_PIXELS, 6, 1,
			u.Health, u.GetStaticData().MaxHitpoints, r.getFactionColor(u.Faction), &rl.Black, &rl.DarkGray)
	}
}

func (r *renderer) renderBuilding(b *Battlefield, pc *playerController, bld *Building) {
	x, y := geometry.TileCoordsToTrueCoords(bld.TopLeftX, bld.TopLeftY)
	x -= 0.5
	y -= 0.5
	osx, osy := r.physicalToOnScreenCoords(x, y)
	w, h := bld.GetStaticData().W, bld.GetStaticData().H
	// log.Printf("%d, %d \n", osx, osy)
	if !r.isRectInViewport(osx, osy, int32(w*TILE_SIZE_IN_PIXELS), int32(h*TILE_SIZE_IN_PIXELS)) {
		return
	}

	var sprites []rl.Texture2D
	sprites = []rl.Texture2D{
		buildingsAtlaces[bld.GetStaticData().SpriteCode].GetSpriteByColorAndFrame(0, 0),
	}
	for _, s := range sprites {
		rl.DrawTexture(
			s,
			osx,
			osy,
			DEFAULT_TINT,
		)
	}
	if bld.Faction != nil && (bld.GetStaticData().W > 1 || bld.GetStaticData().H > 1) {
		r.renderFactionFlagAt(
			b,
			bld.Faction,
			osx+4,
			osy+int32(bld.GetStaticData().H*TILE_SIZE_IN_PIXELS),
		)
	}
	if !bld.IsFullyCaptured() {
		r.renderBuildingCaptureUI(bld, osx+TILE_SIZE_IN_PIXELS/4, osy+int32(bld.GetStaticData().H*TILE_SIZE_IN_PIXELS)-24)
	}
}

func (r *renderer) renderBuildingCaptureUI(bld *Building, x, y int32) {
	r.drawOutlinedRect(x, y, 18*6, 24, 3, rl.Gray, rl.Black)
	r.drawText("[", x+2, y, 24, rl.White)
	r.drawText("]", x+2+18*5, y, 24, rl.White)
	for i := range bld.CaptureProgress {
		color := rl.Gray
		if bld.CaptureProgress[i] != nil {
			color = spritesatlas.FactionColors[bld.CaptureProgress[i].ColorNumber]
			r.drawText("*", x+int32(18*(i+1)), y, 24, color)
		}
	}
}

func (r *renderer) renderProjectile(proj *Projectile) {
	if proj.SetToRemove {
		return
	}
	x, y := proj.CenterX, proj.CenterY
	osx, osy := r.physicalToOnScreenCoords(x, y)
	// log.Printf("%d, %d \n", osx, osy)
	if !r.isRectInViewport(osx, osy, TILE_SIZE_IN_PIXELS, TILE_SIZE_IN_PIXELS) {
		return
	}
	sprite := projectilesAtlaces[proj.GetStaticData().SpriteCode].GetSpriteByColorDegreeAndFrameNumber(0, proj.RotationDegree, 0)
	rl.DrawTexture(
		sprite,
		osx-sprite.Width/2,
		osy-sprite.Height/2,
		DEFAULT_TINT, // proj.faction.factionColor,
	)
}

func (r *renderer) renderEffect(btl *Battlefield, e *Effect) {
	// debugWritef("Percent is %d", e.getExpirationPercent(r.btl.currentTick))
	if e.CreationTick <= r.btl.CurrentTick && e.GetExpirationPercent(r.btl.CurrentTick) <= 100 {
		x, y := e.CenterX, e.CenterY
		osx, osy := r.physicalToOnScreenCoords(x, y)
		// log.Printf("%d, %d \n", osx, osy)
		if !r.isRectInViewport(osx, osy, TILE_SIZE_IN_PIXELS, TILE_SIZE_IN_PIXELS) {
			return
		}
		neededAtlas := effectsAtlaces[e.GetStaticData().SpriteCode]
		expPercent := e.GetExpirationPercent(r.btl.CurrentTick)
		currentFrame := geometry.GetPartitionIndex(expPercent, 0, 100, neededAtlas.TotalFrames())
		if e.SplashCircleRadius > 0 {
			radius := float32(float64(expPercent*TILE_SIZE_IN_PIXELS)*e.SplashCircleRadius) / 100.0
			rl.DrawCircleLines(osx, osy, radius, rl.Red)
			rl.DrawCircleLines(osx, osy, radius+1, rl.Maroon)
			rl.DrawCircleLines(osx, osy, radius+2, rl.Yellow)
		}
		rl.DrawTexture(
			neededAtlas.GetSpriteByFrame(currentFrame),
			osx-neededAtlas.GetSpriteByFrame(currentFrame).Width/2,
			osy-neededAtlas.GetSpriteByFrame(currentFrame).Height/2,
			DEFAULT_TINT, // proj.faction.factionColor,
		)
	}
}
