package rule

var (
	moveXs = []int{-1, 0, 1};
	moveYs = []int{-1, 0, 1};
)


func CountAliveNeighbor(grid *[][]int, posX int, posY int) (count int) {
	count = 0
	row := len(*grid)
	col := len((*grid)[0])

	for _, moveX := range moveXs {
		for _, moveY := range moveYs {
			curX := posX + moveX
			curY := posY + moveY

			if curX == posX && curY == posY { // posisi sendiri
				continue
			}

			if OutOfGrid(curX, curY, row, col) {
				continue
			}

			if (*grid)[curX][curY] == 1 { count++ }
		}
	}

	return count
}