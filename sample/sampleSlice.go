package sample

import (
	"fmt"
	"github.com/tfajarama/game-of-life/utils/generateboard"
	"github.com/tfajarama/game-of-life/utils/print"
)

var (
	slice = [][]int{{1, 0, 1}, {0, 1, 0}, {0, 0, 1}} // case 1
	//slice = [][]int{{0, 1, 0}, {1, 0, 0}, {0, 1, 0}}	// case 2
	//slice = [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}} // case 3
	banyakIterasi = 3
)

func SliceTwoDimension() {
	print.PrintBoard(&slice)

	for k := 0; k < banyakIterasi; k++ {
		slice = generateboard.NextBoard(&slice)
		fmt.Printf("Generation Number: #%d\n", k+1)
		print.PrintBoard(&slice)
	}
}
