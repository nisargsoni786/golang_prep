package main

import "fmt"
// import "github.com/pkg/errors"

func main() {
	var a int64=10
	var b int32=5
	var mean float64

	mean=(float64(a)+float64(b))/2

	fmt.Println(mean)
}