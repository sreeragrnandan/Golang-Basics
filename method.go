package main

import "fmt"

func MainMethod() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
