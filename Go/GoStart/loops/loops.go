package loops

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SimpleFor(a *int) {
	for i := *a + 1; i < *a+11; i++ {
		fmt.Println("i -> ", i)
	}

	fmt.Println("value of a at last ", *a)
	*a = 100
}

func OnlyFor(a *int) {
	i := 0

	for i <= *a {
		fmt.Println("i -> ", i)
		i += 1
	}

	fmt.Println("value of a at last ", *a)
	*a = 200
}

func LikeWhile(a *int) {
	i := 0

	for {
		if i == *a {
			break
		} else {
			fmt.Println(i)
			i += 1
		}
	}
	fmt.Println("value of a at last ", *a)
	*a = 300
}

func Array() {
	var n int = 11
	arr := make([]primitive.ObjectID, n)
	arr[3] = primitive.NewObjectID()

	fmt.Println("EMMPPPTTTYYYY", arr[2] == primitive.NilObjectID)
	fmt.Println("=======", arr[3])
	fmt.Println("array[2] array[3] -> ", arr[2], arr[3])
	// fmt.Println("array -> ", arr)
}

func Array2() {
	var arr []string

	arr = append(arr, "a")
	arr = append(arr, "b")
	arr = append(arr, "c")

	fmt.Println(arr)

	arr = append(arr[0:1], arr[2:]...)

	fmt.Println(arr)
}

func Array3() {
	type AB struct {
		a int
		b string
	}

	var arr []AB

	arr = append(arr, AB{a: 1, b: "a"})
	arr = append(arr, AB{a: 2, b: "b"})
	arr = append(arr, AB{a: 3, b: "c"})
	arr = append(arr, AB{a: 4, b: "d"})

	fmt.Println(arr)

	// newarr := append(arr[0:1], arr[2:]...)

	// fmt.Println(newarr)

	for i := 7; i < 7; i++ {
		fmt.Println("Yes")
	}

	// for _, v := range arr {
	// 	if v.a > 2 {
	// 		v.a = v.a - 1
	// 	}
	// }
	// arr = append(arr[0:0], arr[0:4]...)
	// arr = append(arr[0:0], arr[2])
	// arr = append(arr, arr[0:0]...)

	fmt.Println(arr)
}

func Array4() {
	type AB struct {
		A int    `json:"a"`
		B string `json:"b"`
		C string `json:"c"`
		D int    `json:"-"`
	}

	a := AB{A: 1, B: "abcd", C: "qsjhfsd", D: 100}
	fmt.Println(reflect.TypeOf(a), a)

	aa, _ := json.Marshal(a)
	fmt.Println(reflect.TypeOf(aa), aa)

	aaa := string(aa)
	fmt.Println(reflect.TypeOf(aaa), aaa)

	var convted AB

	_ = json.Unmarshal([]byte(aaa), &convted)
	fmt.Println(reflect.TypeOf(convted), convted)
	// _ = json.Unmarshal(aa, &convted)
	// fmt.Println(reflect.TypeOf(convted), convted)

}

func Array5() {

	id := primitive.NewObjectID()
	fmt.Println(id)

	sep := "|"

	a := fmt.Sprintf("%s%s%s", "abcd", sep, id.Hex())
	fmt.Println(a)

	dur, _ := time.ParseDuration("10s")
	dur2 := time.Second * 10
	fmt.Println(dur, dur2)
}

func ArrayAppend() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a)

	c := make([]int, len(a))
	copy(c, a)
	// var c []int
	// c = append(c, a...)
	// c = append(c, 1000)

	// var b []int

	// b = append(a[1:3], 100)

	// fmt.Println(a, b, c)

	// b := append(a[:3], 300)
	// fmt.Println(a)

	// a = append(a, c[3:]...)
	// fmt.Println(a, b)

	a = append(a[:3+1], a[3:]...)
	fmt.Println(a)

}
