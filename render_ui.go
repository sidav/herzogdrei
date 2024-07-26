package main

import (
	"fmt"
	"herzog/game"
	"herzog/game/game_static"
	"herzog/lib/geometry"
	"herzog/lib/strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *renderer) renderUI(b *game.Battlefield, f *game.Faction) {
	r.drawStatusbars(b, f)
	if f.Commander.CarriedUnit != nil {
		r.drawPickedUpUnitInfo(f, WINDOW_W-320, 0)
	} else {
		if f.ProductionInProgress() {
			r.drawBuildProgressBar(f)
		} else {
			r.drawBuildBar(f)
		}
	}
	u := b.GetGroundUnitAtTileCoordinates(geometry.TrueCoordsToTileCoords(f.Commander.GetPhysicalCenterCoords()))
	if u != nil {
		r.drawUnitInfo(u, "Unit on ground", WINDOW_W/2-160, WINDOW_H-160)
	}
	if b.CurrentTick%30 == 0 {
		r.updateMinimap(b, f)
	}
	r.renderMinimap(b, f,
		WINDOW_W-(MINIMAP_PIXEL_SIZE*int32(len(r.minimap)))-4,
		WINDOW_H-(MINIMAP_PIXEL_SIZE*int32(len(r.minimap[0])))-4)
}

func (r *renderer) drawStatusbars(b *game.Battlefield, f *game.Faction) {
	const leftXCoord = 4
	r.drawProgressBar(leftXCoord, 1, 344, 28, 4, f.HQBuilding.Hitpoints, f.HQBuilding.GetStaticData().MaxHitpoints, &rl.DarkBlue, &rl.Red, &rl.Black)
	r.drawTextCenteredAt(fmt.Sprintf("HQ STATUS %d/%d", f.HQBuilding.Hitpoints, f.HQBuilding.GetStaticData().MaxHitpoints), 172, 4, 28, rl.White)
	r.drawOutlinedRect(leftXCoord, 32, 128, 36, 3, rl.Red, rl.Black)
	rl.DrawText(fmt.Sprintf("%5d$", f.Gold), leftXCoord, 34, 32, rl.White)

	var textSize int32 = 24
	var height int32 = WINDOW_H - (textSize+8)*3
	height += textSize + 4
	r.drawProgressBar(leftXCoord, height, 344, textSize, 4, f.Commander.AsUnit.Health, f.Commander.GetStaticData().MaxHitpoints, &rl.DarkBlue, &rl.Red, &rl.Black)
	r.drawTextCenteredAt(fmt.Sprintf("HEALTH %d/%d", f.Commander.AsUnit.Health, f.Commander.GetStaticData().MaxHitpoints), 172, height, textSize, rl.White)
	height += textSize + 4
	r.drawProgressBar(leftXCoord, height, 344, textSize, 4, f.Commander.AsUnit.Fuel, f.Commander.GetStaticData().MaxFuel, &rl.DarkBlue, &rl.Red, &rl.Black)
	r.drawTextCenteredAt(fmt.Sprintf("SOLYARKA %d/%d", f.Commander.AsUnit.Fuel, f.Commander.GetStaticData().MaxFuel), 172, height, textSize, rl.White)
}

func (r *renderer) drawBuildBar(f *game.Faction) {
	code, order := f.GetSelectedProduction()
	static := game_static.GetUnitStaticDataByCode(code)
	r.drawOutlinedRect(WINDOW_W-320, 0, 320, 140, 3, rl.Red, rl.Black)
	r.drawTextCenteredAt("BUILD", WINDOW_W-160, 3, 32, rl.White)
	r.drawTextCenteredAt(strings.CenterStringBetween("<Q", static.DisplayedName, ">E", 16),
		WINDOW_W-160, 36, 32, rl.Beige)
	orderStr := fmt.Sprintf("%s -$%d", game.GetOrderName(order), static.OrderCosts[order])
	r.drawTextCenteredAt(strings.CenterStringBetween("<A", orderStr, ">D", 20),
		WINDOW_W-160, 68, 28, rl.Orange)
	r.drawTextCenteredAt(fmt.Sprintf("TOTAL $%d", f.GetTotalCostForCurrentProduction()),
		WINDOW_W-160, 100, 32, rl.White)
}

func (r *renderer) drawBuildProgressBar(f *game.Faction) {
	const windowWidth = 320
	r.drawOutlinedRect(WINDOW_W-windowWidth, 0, windowWidth, 140, 3, rl.Red, rl.Black)
	percentage := f.GetCurrentProductionPercentage()
	if percentage != 100 {
		r.drawTextCenteredAt(fmt.Sprintf("BUILDING %d%%", percentage), WINDOW_W-windowWidth/2, 3, 32, rl.White)
		r.drawProgressBar(WINDOW_W-windowWidth, 124, windowWidth, 16, 4,
			f.GetCurrentProductionPercentage(), 100,
			&rl.DarkGreen, &rl.Black, &rl.Red)
	} else {
		r.drawTextCenteredAt("READY", WINDOW_W-windowWidth/2, 3, 32, rl.White)
	}
	code, order := f.GetSelectedProduction()
	r.drawTextCenteredAt(game_static.GetUnitStaticDataByCode(code).DisplayedName, WINDOW_W-windowWidth/2, 40, 32, rl.White)
	r.drawTextCenteredAt(game.GetOrderName(order), WINDOW_W-windowWidth/2, 72, 32, rl.White)
}

func (r *renderer) drawPickedUpUnitInfo(f *game.Faction, x, y int32) {
	const windowWidth = 320
	unit := f.Commander.CarriedUnit
	unitOrderCode := unit.Order.Code
	expectedOrderCode := f.GetOrderToBeGivenOnDrop()

	r.drawOutlinedRect(x, y, windowWidth, 160, 3, rl.Red, rl.Black)
	r.drawTextCenteredAt("Carried unit", x+windowWidth/2, y+4, 32, rl.White)
	r.drawTextCenteredAt(unit.GetStaticData().DisplayedName, x+windowWidth/2, y+36, 32, rl.White)
	var orderStr string
	if unitOrderCode != expectedOrderCode {
		orderStr = fmt.Sprintf("%s -$%d", game.GetOrderName(expectedOrderCode), unit.GetStaticData().OrderCosts[expectedOrderCode])
	} else {
		orderStr = game.GetOrderName(unitOrderCode)
	}
	r.drawTextCenteredAt(strings.CenterStringBetween("<A", orderStr, ">D", 20), x+windowWidth/2, y+72, 28, rl.Beige)
}

func (r *renderer) drawUnitInfo(u *game.Unit, title string, x, y int32) {
	const windowWidth = 320
	r.drawOutlinedRect(x, y, windowWidth, 160, 3, rl.Red, rl.Black)
	r.drawTextCenteredAt(title, x+windowWidth/2, y+4, 32, rl.White)
	r.drawTextCenteredAt(u.GetStaticData().DisplayedName, x+windowWidth/2, y+36, 32, rl.White)
	r.drawTextCenteredAt(game.GetOrderName(u.Order.Code), x+windowWidth/2, y+72, 32, rl.White)
}
