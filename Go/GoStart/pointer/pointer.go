package pointer

import "fmt"

func Pointer() {
	a := 5
	fmt.Println(a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)
}
