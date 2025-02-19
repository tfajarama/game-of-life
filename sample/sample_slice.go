package sample

import "fmt"

func SliceTwoDimension() {

	slice := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			fmt.Println(slice[i][j], "\t")
		}
		fmt.Println()
	}
}
