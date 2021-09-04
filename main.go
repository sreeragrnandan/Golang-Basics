package main

import (
	"fmt"
	"strconv"
)

var p float32 = 30

func main() {
	var i int
	i = 42
	var j int = 27
	k := 99
	fmt.Println(i, j, k, p)
	fmt.Printf("%v, %T \n", j, j)
	// Type Convertion
	var x float32 = 100.99
	xInt := int(x)
	fmt.Printf("%v, %T \n", xInt, xInt)
	xStr := strconv.Itoa(xInt)
	fmt.Printf("%v, %T \n", xStr, xStr)
}
