package main

import "fmt"

func main() {
	map1:=make(map[string]int)

	map1["a"] = 1
	map1["b"] = 2

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	val,ok:=map1["a"]
	fmt.Println(val, ok)

	val,ok=map1["not_there"]
	fmt.Println(val,ok)

}