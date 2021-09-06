package main

import "fmt"

func main() {
	// Varialble types
	Varialble()

	// Constant types
	Const()

	//Array and Slice
	ArraySlice()

	//Map & Structs
	mapStruct()

	// If Else
	fmt.Println("\nIfElse")
	ifElse()

	//For Loop
	fmt.Println("\nForLoop")
	ForLoop()

	// fmt.Println("\nPanicFun")
	// PanicFun()

	//Function
	for i := 0; i < 5; i++ {
		sayMessage("Hello Go!", i)
	}
	// Variatic parameters
	Sum(1, 2, 3, 4, 5, 6, 7, 8)

	//Pass by reference
	greetings := "Hello"
	name := "Shilpa"
	SayGreeting(&greetings, &name)
	fmt.Println(name)

	//Returning Function
	res := retInt(2)
	fmt.Println("retInt(2):", res)

	//Returning via finction signature
	res2 := retInt2(2)
	fmt.Println("retInt2(2):", res2)

	//Result with error
	d, err := divide(2, 0)
	fmt.Println("divide(2, 0)", d, err)

	// Method
	MainMethod()
	// Interface
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Interface Hello Go!"))
}
