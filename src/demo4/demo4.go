package main

import "fmt"

func testselect() {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <-ch
		fmt.Println("receive value=", i)
	}
}

func testselect1() {
	ch := make(chan int, 1)

	select {
	case ch <- 0:
	case ch <- 1:
	case ch <- 2:
	case ch <- 3:
	}
	i := <-ch
	fmt.Println("receive value=", i)

}

func main() {
	//testselect()

	testselect1()
}
