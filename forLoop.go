package main

import "fmt"

func ForLoop() {
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum) // 10 (1+2+3+4)

	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n) // 8 (1*2*2*2)

	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	slice := make([]int, 10, 100)
	fmt.Println(slice)
	for i := 0; i < 10; i++ {
		slice[i] = i
	}

	fmt.Println(slice)

}
