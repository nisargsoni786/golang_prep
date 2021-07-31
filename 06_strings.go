package main

import (
	"fmt"
)

func main() {
	a := "nisarg"

	for _, i := range a {
		fmt.Println(i)
		chr := string(i)
		fmt.Println((chr))
	}
}
