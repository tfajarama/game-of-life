package print

import "fmt"

func PrintBoard(slice *[][]int) {
	for i := 0; i < len((*slice)); i++ {
		for j := 0; j < len((*slice)[i]); j++ {
			fmt.Print((*slice)[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}