package main

import (
	"fmt"
)

var a string = "abc"
var b string = "def"

var a1 uint = 10
var a2 float32 = 20000000000000000000084770000000.45
var c1 byte = 'x'

func main() {
	c2 := 'b'
	// a := 100012
	fmt.Println(a+b, "ds")
	fmt.Println(a1, a2)
	fmt.Printf("c1 %c and c2 %c", c1, c2)
}
