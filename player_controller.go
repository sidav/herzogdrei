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

	if !pc.faction.ProductionInProgress() && com.CarriedUnit == nil {
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
			pc.faction.StartProduction()
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
	code, order := pc.faction.GetSelectedProduction()
	totalListSize := len(game_static.STableUnits)
	changed := false
	for !changed || game_static.STableUnits[code].IsCommander {
		changed = true
		code += step
		if code >= totalListSize {
			code = 0
		}
		if code < 0 {
			code += totalListSize
		}
	}
	// Select next available order
	for !game_static.STableUnits[code].CanDoOrder(order) {
		order = (order + 1) % game_static.ORDERS_TOTAL
	}
	pc.faction.SetSelectedProduction(code, order)
}

func (pc *playerController) selectNextOrderForBuildable(step int) {
	if step == 0 {
		panic("Wat")
	}
	code, order := pc.faction.GetSelectedProduction()
	totalListSize := game_static.ORDERS_TOTAL
	done := false
	for !done {
		order += step
		if order >= totalListSize {
			order = 0
		}
		if order < 0 {
			order += totalListSize
		}
		done = game_static.STableUnits[code].CanDoOrder(order)
	}
	pc.faction.SetSelectedProduction(code, order)
}
