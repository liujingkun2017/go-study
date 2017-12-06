package main

import (
	"config"
	d1 "demo1"
	d2 "demo2"
	d3 "demo3"
	"fmt"
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

	fmt.Println("demo1---------------------------------")
	fmt.Println(d1.Integer(1).Less(2))

	var a Integer = 1
	var b Integer = 1

	a.test(b)
	a.test1(b)
	a.test(b)
	a.test2(b)
	a.test(b)

	fmt.Println("demo2---------------------------------")

	for i := 1; i < 10; i++ {
		go d2.Add(i, i)
	}

	fmt.Println("demo3---------------------------------")
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go d3.Count(chs[i])
	}
	for _, ch := range chs {
		x := <-ch
		fmt.Println("x=", x)
	}

}
