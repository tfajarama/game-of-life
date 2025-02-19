package rule

func OutOfGrid(posX int, posY int, row int, col int) bool {
	return (posX < 0 ||
		posY < 0 ||
		posX >= row ||
		posY >= col)
}
