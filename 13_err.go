package main

import ("fmt"
		"reflect"
		"log"
)

func dev(a,b int) float64{
	if b==0{
		panic("Zeroooo")
	}
	return float64(a)/float64(b)
}

func main() {
	defer func() {
		if err := recover();err != nil{
			log.Println("errorrr........................",err)
		}
	}()
	ans:=dev(4,2)
	fmt.Println(ans)
}