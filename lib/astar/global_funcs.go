package astar

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getCellWithCoordsFromList(list *[]*Cell, x, y int) *Cell {
	for _, c := range *list {
		if c.X == x && c.Y == y {
			return c
		}
	}
	return nil
}

func getCellWithLowestHeuristicFromList(list *[]*Cell) *Cell {
	lowest := (*list)[0]
	for _, c := range *list {
		if c.h < lowest.h {
			lowest = c
		}
	}
	return lowest
}
