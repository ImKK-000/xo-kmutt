package game

type Board struct {
	Grid [][]string
	Size int
}

func NewBoard(size int) Board {
	grid := make([][]string, size)
	for index := range grid {
		grid[index] = make([]string, size)
	}
	return Board{
		Size: size,
		Grid: grid,
	}
}
