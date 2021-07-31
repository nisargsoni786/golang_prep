package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 5
	bf := 5.0
	var c float64
	c = float64(b) / float64(a)

	fmt.Println(a + b)
	fmt.Println(c)

	if float64(b) == bf {
		fmt.Println("YES")
	}

	fmt.Println(a << 1)
}
