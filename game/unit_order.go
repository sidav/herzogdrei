package game

import (
	. "herzog/game/game_static"
)

type Order struct {
	Code             int
	TargetX, TargetY int // tile coords
	TargetActor      Actor
	OriginX, OriginY int // Origin coords (where the order was received) of an order, are order-dependent in use
}

func (o *Order) Reset() {
	o.Code = ORDER_STANDBY
	o.TargetX = -1
	o.TargetY = -1
	o.TargetActor = nil
}

func (o *Order) ResetTargets() {
	o.TargetX = -1
	o.TargetY = -1
	o.TargetActor = nil
}

func (o *Order) SetTargetTile(x, y int) {
	o.TargetX = x
	o.TargetY = y
}

func (o *Order) SetOrigin(x, y int) {
	o.OriginX = x
	o.OriginY = y
}

func (o *Order) GetTargetTileCoords() (int, int) {
	return o.TargetX, o.TargetY
}

func GetOrderName(code int) string {
	switch code {
	case ORDER_STANDBY:
		return "Stand by"
	case ORDER_PATROL:
		return "Patrol"
	case ORDER_SEARCHNDESTROY:
		return "Hunt"
	case ORDER_CAPTURE:
		return "Capture"
	}
	panic("No order name")
}
