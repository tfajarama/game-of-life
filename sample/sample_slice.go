package sample

import (
	"fmt"

	"github.com/tfajarama/game-of-life/utils/generateboard"
	"github.com/tfajarama/game-of-life/utils/print"
)

func SliceTwoDimension() {

	slice := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	print.PrintBoard(&slice)

	banyakIterasi := 3

	for k := 0; k < banyakIterasi; k++ {
		slice = generateboard.NextBoard(&slice)
		fmt.Printf("Generation Number: #%d\n", k+1)
		print.PrintBoard(&slice)
	}
}
