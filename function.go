package main

import (
	"fmt"
)

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

func Sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is", result)
}

func SayGreeting(greetings, name *string) {
	*name = "Sreerag"
	fmt.Println(*greetings, *name)
}

func retInt(k int) int {
	ans := k + 1
	return ans
}

func retInt2(k int) (ans int) {
	ans = k + 1
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divided by zero")
	}
	return a / b, nil
}
