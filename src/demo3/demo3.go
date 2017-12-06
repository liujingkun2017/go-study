package demo3

import "fmt"

func Count(ch chan int) {
	ch <- 1
	fmt.Println("counting")
}
