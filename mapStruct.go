package main

import "fmt"

func Map() {
	countryCode := make(map[string]int)
	countryCode = map[string]int{
		"India": 91,
		"US":    1,
		"UK":    44,
	}
	fmt.Println(countryCode)
	countryCode["Dubai"] = 971
	fmt.Println(countryCode)
	delete(countryCode, "Dubai")
	fmt.Println(countryCode)
	val, ok := countryCode["Inda"]
	fmt.Println(val, ok)
	val, ok = countryCode["India"]
	fmt.Println(val, ok)
}

func normalStruct() {
	type Student struct {
		RollNumber    int
		AvgMark       int
		Name          string
		subjectPassed []string
	}

	aStudent := Student{
		RollNumber: 3,
		Name:       "Sreerag",
		subjectPassed: []string{
			"ALL", "Sub1", "Sub2",
		},
	}
	fmt.Println(aStudent.Name)
	fmt.Println(aStudent)
}

func mapStruct() {
	fmt.Println("\n\nMap")
	Map()

	fmt.Println("\n\nStruct")
	normalStruct()

}
