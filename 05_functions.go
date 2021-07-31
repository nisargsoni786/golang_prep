package main

import (
	"fmt"
)

func abc(a int, b int) int {
	return a + b
}

func multi(x, y string) (string, string) {
	return y, x
}

func main() {
	ans := abc(15, 5)
	fmt.Println(ans)

	a1, a2 := multi("abc", "def")
	fmt.Println(a1, a2)

}
