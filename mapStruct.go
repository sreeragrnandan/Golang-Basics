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

func mapStruct() {
	fmt.Println("\n\nMap")
	Map()
}
