package main3

import (
	"fmt"

	"github.com/nisarg/stringutil"
)

const (
	A = iota
	B = iota * 100 * iota
	C = iota * 200 * iota
	D = iota * 300 * iota
)

func Abc() string {
	fmt.Println(A, B, C, D)
	fmt.Println(stringutil.Reversec("! ZZZ oG ,olleH"))
	fmt.Println(stringutil.Reverse("! ZZZ oG ,olleH"))
	// fmt.Println(stringutil.reverseTwo("! ZZZ oG ,ollsdfvbsdfgdfgeH"))
	return "Abc Completed"
}
