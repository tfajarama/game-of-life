package generateboard

import (
	"github.com/tfajarama/game-of-life/utils/rule"
)

func NextBoard(grid *[][]int) (nextGrid [][]int) {
	row := len(*grid)
	col := len((*grid)[0])

	nextGrid = make([][]int, row)
	for i := 0; i < row; i++ {
		nextGrid[i] = make([]int, col)
	}

	for i := 0; i < len((*grid)); i++ {
		for j := 0; j < len((*grid)[i]); j++ {
			nextGrid[i][j] = rule.SetLiveOrDie((*grid)[i][j], rule.CountAliveNeighbor(&(*grid), i, j))
		}
	}
	return nextGrid
}