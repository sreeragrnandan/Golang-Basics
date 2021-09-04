package main

import "fmt"

func Array() {
	fmt.Println("\n Array")
	// Static
	var a [3]int
	fmt.Println(a)
	grades := [3]int{97, 85, 72}
	fmt.Println("Grades", grades)

	grades2 := [...]int{97, 85, 72}
	fmt.Println("Grades2", grades2)

	//String
	var students [3]string
	fmt.Printf("Students: %v\n", students)

	students[0] = "Lisa"
	students[1] = "Arnold"
	students[2] = "Aline"

	fmt.Printf("Students: %v\n", students)

	fmt.Printf("Number of Students: %v\n", len(students))
	// Matrix
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	// Slice
	sliceA := []int{1, 2, 3, 4}
	sliceB := make([]int, 10, 100)
	sliceA[0] = 100
	sliceB = append(sliceB, 2)
	sliceB = append(sliceB, sliceA...)

	fmt.Printf("sliceA %v \nsliceB %v \n", sliceA, sliceB)
	fmt.Println("sliceA[1:2]", sliceA[1:2], "Length", len(sliceA), "Capacity", cap(sliceA))
	fmt.Println("sliceB[1:2]", sliceA[1:2], "Length", len(sliceB), "Capacity", cap(sliceB))
}
