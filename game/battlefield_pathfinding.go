package game

import "herzog/lib/astar"

func (b *Battlefield) getVectorForPath(fromX, fromY, targetX, targetY int) (int, int) {
	if b.pathfinder == nil {
		b.pathfinder = &astar.AStarPathfinder{
			MapWidth:            len(b.Tiles),
			MapHeight:           len(b.Tiles[0]),
			ForceGetPath:        true,
			DiagonalMoveAllowed: true,
		}
	}
	b.pathfinder.ForceIncludeFinish = false

	return b.pathfinder.FindPath(
		func(tx, ty int) int {
			if b.AreCoordsPassable(tx, ty) {
				return 10
			}
			return -1
		},
		fromX, fromY, targetX, targetY,
	).GetNextStepVector()
}

// func (b *Battlefield) getVectorForForcedPath(fromX, fromY, targetX, targetY int) (int, int) {
// 	b.pathfinder.ForceIncludeFinish = true
// 	return b.getVectorForPath(fromX, fromY, targetX, targetY)
// }
