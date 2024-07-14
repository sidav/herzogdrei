package game

import (
	"herzog/lib/geometry"
	"math"
)

func (b *Battlefield) ExecuteUnitAction(u *Unit) {
	switch u.Action.Kind {
	case ACTION_NONE:
		return
	case ACTION_ROTATE:
		b.ExecuteRotateToVectorActionForUnit(u)
	case ACTION_START_MOVING:
		b.ExecuteStartMovingActionForUnit(u)
	case ACTION_MOVE:
		b.ExecuteMoveActionForUnit(u)
	case ACTION_ENTER_BUILDING:
		b.ExecuteEnterBuildingActionForUnit(u)
	}
}

// Returns true if rotation occured
func (b *Battlefield) ExecuteRotateToVectorActionForUnit(u *Unit) {
	if b.TryRotateToVectorForUnit(u) {
		// do nothing?
	} else {
		u.Action.Reset()
	}
}

// Rotate and take care of collision detection here
func (b *Battlefield) ExecuteStartMovingActionForUnit(u *Unit) {
	// Is it needed to rotate?
	if b.TryRotateToVectorForUnit(u) {
		return
	}
	vx, vy := u.Action.Vx, u.Action.Vy
	utx, uty := u.GetTileCoords()
	nextTileX, nextTileY := utx+int(math.Round(vx)), uty+int(math.Round(vy))
	if !b.AreCoordsPassable(nextTileX, nextTileY) {
		u.Action.Reset()
		return
	}

	b.Tiles[utx][uty].landActorHere = nil
	b.Tiles[nextTileX][nextTileY].landActorHere = u
	u.Action.Kind = ACTION_MOVE
}

// Movement itself is here
func (b *Battlefield) ExecuteMoveActionForUnit(u *Unit) {
	moveSpeed := u.GetStaticData().MovementSpeed
	vx, vy := u.Action.Vx, u.Action.Vy
	u.CenterX += vx * moveSpeed
	u.CenterY += vy * moveSpeed
	u.Action.MovementRemaining -= moveSpeed

	// cx, cy := u.GetPhysicalCenterCoords()
	tx, ty := u.GetTileCoords()
	tileCenterX, tileCenterY := geometry.TileCoordsToTrueCoords(tx, ty)
	// Stop movement if needed
	if u.Action.MovementRemaining < 0 {
		u.Action.Reset()
		// Snap to tile center
		u.CenterX = tileCenterX
		u.CenterY = tileCenterY
	}
}

// Rotate and the enter itself are here
func (b *Battlefield) ExecuteEnterBuildingActionForUnit(u *Unit) {
	if b.TryRotateToVectorForUnit(u) {
		return
	}

	// Start entering. Remove the unit from tile
	tx, ty := u.GetTileCoords()
	if b.GetGroundActorAtTileCoordinates(tx, ty) == u {
		b.Tiles[tx][ty].landActorHere = nil
	}
	// Enter the building
	moveSpeed := u.GetStaticData().MovementSpeed
	vx, vy := u.Action.Vx, u.Action.Vy
	u.CenterX += vx * moveSpeed
	u.CenterY += vy * moveSpeed
	u.Action.MovementRemaining -= moveSpeed
	// The building is entered
	if u.Action.MovementRemaining < 0.005 {
		b.removeActorFromList(u)
		b.progressBuildingCapture(u.Order.TargetActor.(*Building), u.Faction)
	}
}

func (b *Battlefield) TryRotateToVectorForUnit(u *Unit) bool {
	requiredAngle := geometry.GetDegreeOfFloatVector(u.Action.Vx, u.Action.Vy)
	if u.ChassisDegree != requiredAngle {
		u.rotateChassisTowardsVector(u.Action.Vx, u.Action.Vy)
		return true
	}
	return false
}
