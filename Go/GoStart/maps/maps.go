package maps

import (
	"fmt"
	"reflect"
)

func SimpleMaps() map[int]string {
	a := map[int]string{
		1: "N",
		2: "I",
		3: "S",
		4: "A",
		5: "R",
		6: "G",
	}
	fmt.Println(a[7] == "")

	return a
}

func DoubleMaps() map[int]map[int]string {
	a := map[int]string{
		1: "N",
		2: "I",
		3: "S",
		4: "A",
		5: "R",
		6: "G",
	}
	b := map[int]string{
		11: "XN",
		12: "XI",
		13: "XS",
		14: "XA",
		15: "XR",
		16: "XG",
	}

	var dm = make(map[int]map[int]string)
	dm[100] = a
	dm[200] = b

	fmt.Println(len(dm[300]))

	return dm
}

func InrterfaceMaps() any {
	a := map[int]string{
		1: "N",
		2: "I",
		3: "S",
		4: "A",
		5: "R",
		6: "G",
	}
	b := map[int]string{
		11: "XN",
		12: "XI",
		13: "XS",
		14: "XA",
		15: "XR",
		16: "XG",
	}

	var dm = make(map[int]any)
	dm[100] = a
	dm[200] = b
	dm[300] = "Nisarg Soni"
	fmt.Println("Type -> ", reflect.TypeOf(dm))
	fmt.Println("dm 300 -> ", dm[300])
	return dm
}
