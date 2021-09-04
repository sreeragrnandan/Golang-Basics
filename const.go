package main

import (
	"fmt"
)

const (
	isAdmin = 1 << iota
	isHeadQuaters
	canSeeFinancials

	canSeeAsia
)

func Const() {
	const i int = 20
	fmt.Println("Constant")
	fmt.Println("Const", i)

	var roles byte = isAdmin | canSeeFinancials | canSeeAsia
	fmt.Printf("%b \n", roles)
	fmt.Printf("Is Admin? %v", isAdmin&roles == isAdmin)
}
