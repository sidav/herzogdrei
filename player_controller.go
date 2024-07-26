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

	if !pc.faction.ProductionInProgress() {
		if com.CarriedUnit == nil {
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
		} else {
			if rl.IsKeyPressed(rl.KeyA) {
				pc.selectNextOrderForPickedUpUnit(-1)
			}
			if rl.IsKeyPressed(rl.KeyD) {
				pc.selectNextOrderForPickedUpUnit(1)
			}
		}
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		if com.CarriedUnit == nil {
			com.SetPickupState()
			return
		} else {
			com.SetDropState()
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
		com.SetMoveState(vx, vy)
	} else {
		com.ResetState()
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

func (pc *playerController) selectNextOrderForPickedUpUnit(step int) {
	if step == 0 {
		panic("Wat")
	}
	static := pc.faction.Commander.CarriedUnit.GetStaticData()
	order := pc.faction.GetOrderToBeGivenOnDrop()
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
		done = static.CanDoOrder(order)
	}
	pc.faction.SetOrderToBeGivenOnDrop(order)
}
