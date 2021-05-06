package world

import "fmt"

var Bounds [][]int

func init() {
	Bounds := [4][4]int{
		{1, 1, 1, 1},
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{1, 1, 1, 1}}

	fmt.Printf("Bounds: %v\n", Bounds)
}
