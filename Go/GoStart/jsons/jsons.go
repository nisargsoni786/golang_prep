package jsons

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Fname string
	Lname string
	Sex   string
}

type Person2 struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Sex   string `json:"sex"`
}

func GetJson() {
	JsonData := `{"fname":"Nisarg","lname":"Soni","sex":"Male"}`
	obj := make(map[string]string)
	json.Unmarshal([]byte(JsonData), &obj)
	fmt.Println("jsondata to Object :- ", reflect.TypeOf(obj), obj)

	dd, _ := json.Marshal(obj)
	fmt.Println("Object to jsondata :- ", reflect.TypeOf(dd), string(dd))

}

func ToMarshal() {
	p := Person{Fname: "Nisarg", Lname: "Soni", Sex: "Male"}
	p2 := Person2{Fname: "Nisarg2", Lname: "Soni2", Sex: "Male2"}
	fmt.Println("p :-", p, "\np2 :- ", p2)

	pm, _ := json.Marshal(p)
	p2m, _ := json.Marshal(p2)

	fmt.Println("pm :-", string(pm), "\np2m :- ", string(p2m))
}
