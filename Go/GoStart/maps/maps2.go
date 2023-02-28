package maps

import "fmt"

type Abc struct {
	Id    string
	Leval string
	Count int
}

func GetMapsAbc() map[string]map[string]int {
	a1 := Abc{Id: "ABC", Leval: "Easy", Count: 2}
	a2 := Abc{Id: "DEF", Leval: "Easy", Count: 4}
	a3 := Abc{Id: "ABC", Leval: "Medium", Count: 5}
	a4 := Abc{Id: "DEF", Leval: "Medium", Count: 4}
	a5 := Abc{Id: "ABC", Leval: "Hard", Count: 2}
	a6 := Abc{Id: "XYZ", Leval: "Easy", Count: 3}
	a7 := Abc{Id: "XYZ", Leval: "Hard", Count: 1}

	a := []Abc{a1, a2, a3, a4, a5, a6, a7}

	anss := make(map[string]map[string]int)
	for _, v := range a {
		Idmap := anss[v.Id]
		if Idmap == nil {
			Idmap = make(map[string]int)
		}
		Idmap[v.Leval] = v.Count
		anss[v.Id] = Idmap
	}

	for k, v := range anss {
		fmt.Println(k, v)
	}
	fmt.Println("\n\n")
	return anss

}

func GetValues(Id string, Leval string) int {
	mapp := GetMapsAbc()

	// InnerMap := mapp[Id]
	// fmt.Println(InnerMap, InnerMap == nil)

	// ans := InnerMap[Leval]
	ans := mapp[Id][Leval]
	return ans

}
