package main

import "fmt"

func abc(a *int) {
	*a = *a + 5
}

func main() {
	a := 10
	fmt.Println(&a)
	abc(&a)
	fmt.Println(a)
}
