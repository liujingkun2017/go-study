package main

import "fmt"

func testbuffer() {

	c := make(chan int, 3)
	for i := 1; i < 5; i++ {
		go w2c(c)
	}

	for i := range c {
		fmt.Println("receivced value ", i)
	}

}

func w2c(c chan int) {
	for {
		select {
		case c <- 0:
			fmt.Println("w2c value 0")
		case c <- 1:
			fmt.Println("w2c value 1")
		case c <- 2:
			fmt.Println("w2c value 2")
		case c <- 3:
			fmt.Println("w2c value 3")
		}
	}

}

func main() {
	testbuffer()
}
