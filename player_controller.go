package main

import (
	"herzog/game"
	"herzog/game/game_static"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type playerController struct {
	faction *game.Faction
}

func (pc *playerController) init(f *game.Faction) {
	pc.faction = f
	pc.selectNextBuildableCode(1)
}

func (pc *playerController) control() {
	com := pc.faction.Commander

	if !pc.faction.BuildsUnitNow && com.CarriedUnit == nil {
		if rl.IsKeyPressed(rl.KeyQ) {
			pc.selectNextBuildableCode(-1)
		}
		if rl.IsKeyPressed(rl.KeyE) {
			pc.selectNextBuildableCode(1)
		}

		if rl.IsKeyPressed(rl.KeyA) {
			pc.selectNextOrderForBuildable(-1)
		}
		if rl.IsKeyPressed(rl.KeyD) {
			pc.selectNextOrderForBuildable(1)
		}

		if rl.IsKeyPressed(rl.KeyEnter) {
			pc.faction.BuildsUnitNow = true
		}
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		if com.CarriedUnit == nil {
			com.AsUnit.Action.Kind = game.ACTION_CPICKUP
			return
		} else {
			com.AsUnit.Action.Kind = game.ACTION_CDROP
			return
		}
	}

	if rl.IsKeyPressed(rl.KeyTab) {
		com.IsTransforming = true
	}

	// Attack
	com.IsFiring = rl.IsKeyDown(rl.KeyLeftControl)
	// Movement
	vx, vy := 0.0, 0.0
	if rl.IsKeyDown(rl.KeyRight) {
		vx = 1.0
	}
	if rl.IsKeyDown(rl.KeyDown) {
		vy = 1.0
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		vx = -1.0
	}
	if rl.IsKeyDown(rl.KeyUp) {
		vy = -1.0
	}
	if vx != 0 || vy != 0 {
		com.AsUnit.Action.Kind = game.ACTION_CMOVE
		com.AsUnit.Action.SetVector(vx, vy)
	} else {
		com.AsUnit.Action.Kind = game.ACTION_NONE
	}
}

func (pc *playerController) selectNextBuildableCode(step int) {
	totalListSize := len(game_static.STableUnits)
	changed := false
	for !changed || game_static.STableUnits[pc.faction.CurrentBuiltUnitCode].IsCommander {
		changed = true
		pc.faction.CurrentBuiltUnitCode += step
		if pc.faction.CurrentBuiltUnitCode >= totalListSize {
			pc.faction.CurrentBuiltUnitCode = 0
		}
		if pc.faction.CurrentBuiltUnitCode < 0 {
			pc.faction.CurrentBuiltUnitCode += totalListSize
		}
	}
	pc.selectNextOrderForBuildable(0)
}

func (pc *playerController) selectNextOrderForBuildable(step int) {
	totalListSize := game_static.ORDERS_TOTAL
	done := false
	for !done {
		pc.faction.CurrentBuiltUnitOrderCode += step
		if pc.faction.CurrentBuiltUnitOrderCode >= totalListSize {
			pc.faction.CurrentBuiltUnitOrderCode = 0
		}
		if pc.faction.CurrentBuiltUnitOrderCode < 0 {
			pc.faction.CurrentBuiltUnitOrderCode += totalListSize
		}
		_, done = game_static.STableUnits[pc.faction.CurrentBuiltUnitCode].OrderCosts[pc.faction.CurrentBuiltUnitOrderCode]
		if step == 0 { // step 0 means "select next if current is not applicable"
			step = 1
		}
	}
}
