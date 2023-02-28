package structures

import "fmt"

type Person struct {
	name   string
	age    int
	gender string
}

type TeamLead struct {
	Technology string
	Person     Person
	Experience int
}

func BasicStruct() {
	p1 := Person{name: "Nisarg Soni", gender: "male"}
	fmt.Println(p1)
}
func getBasicStruct() Person {
	p1 := Person{name: "Nisarg Soni", age: 23, gender: "male"}
	return p1
}

func ComplexStruct() {
	person := getBasicStruct()
	fmt.Println(person)
	teamlead := TeamLead{Technology: "python-Go", Experience: 10, Person: person}
	fmt.Println(teamlead.Technology)
}
