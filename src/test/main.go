package main

import (
	"config"
	"fmt"
	d "demo1"
)

type Integer int

func (a Integer) test(b Integer) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
}

func (a *Integer) test1(b Integer) {
	*a = *a + b
}

func (a Integer) test2(b Integer) {
	a = a + b
}

func main() {
	config.LoadConfig()

	fmt.Println("Hello, GO!")
	fmt.Println(d.Integer(1).Less(2))

	var a Integer = 1
	var b Integer = 1

	a.test(b)

	a.test1(b)

	a.test(b)

	a.test2(b)

	a.test(b)

}
