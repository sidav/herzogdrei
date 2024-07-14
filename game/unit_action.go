package game

import (
	"herzog/lib/geometry"
)

const (
	ACTION_NONE = iota
	// Actions with "C" are commander-only
	ACTION_CMOVE
	ACTION_CPICKUP
	ACTION_CDROP
	// Unit-only actions
	ACTION_ROTATE
	ACTION_START_MOVING
	ACTION_MOVE
	ACTION_ENTER_BUILDING
)

type UnitAction struct {
	Kind              int
	Vx, Vy            float64
	MovementRemaining float64
}

func (a *UnitAction) Reset() {
	a.Kind = ACTION_NONE
	a.MovementRemaining = 0
}

func (a *UnitAction) CalcMoveRemaining() {
	a.MovementRemaining = 1
	if a.Vx*a.Vy != 0 {
		a.MovementRemaining = 1.4
	}
}

func (a *UnitAction) SetVectorByInt(vx, vy int) {
	a.Vx, a.Vy = geometry.VectorToUnitVectorFloat64(float64(vx), float64(vy))
	// a.Vx = float64(vx)
	// a.Vy = float64(vy)
}

func (a *UnitAction) SetVector(vx, vy float64) {
	a.Vx = vx
	a.Vy = vy
}
