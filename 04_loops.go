package main

import (
	"fmt"
)

var x int = 100
var y int = 95

func main() {

	arr := [6]int{11, 12, 13, 14, 15, 16}

	for i := 1; i < 6; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	for i, j := range arr {
		fmt.Println(i, j)
	}

	for x > y { // while loop by for
		fmt.Println(y)
		y++
	}

}
