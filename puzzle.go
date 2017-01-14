package main

import "fmt"

type Test struct {
	W     int
	H     int
	slots [][]int
}

func main() {
	matrix := make([][]int, 4)
	for i := range matrix {
		matrix[i] = make([]int, 4)
	}
	t := Test{
		W:     4,
		H:     4,
		slots: matrix,
	}

	fmt.Printf("%#v\n", t)
}
