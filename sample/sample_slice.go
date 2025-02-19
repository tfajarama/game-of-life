package sample

import (
	"fmt"
	"github.com/tfajarama/game-of-life/utils/rule"
)

func SliceTwoDimension() {

	slice := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			fmt.Print(slice[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			fmt.Print(rule.CountAliveNeighbor(&slice, i, j), " ")
		}
		fmt.Println()
	}
}
