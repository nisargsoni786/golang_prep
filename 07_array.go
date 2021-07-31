package main

import "fmt"

var a = []int{11, 12, 13}

func getarr(arr []int) int {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	return sum
}

func main() {
	for index, i := range a {
		fmt.Println(index, i)
	}

	fmt.Println(a[1])

	fmt.Println(getarr(a))
}
