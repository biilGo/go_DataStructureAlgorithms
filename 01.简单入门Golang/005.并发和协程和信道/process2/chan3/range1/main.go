package main

import "fmt"

func main() {
	buffedChan := make(chan int, 2)
	buffedChan <- 2
	buffedChan <- 3

	for i := range buffedChan {
		fmt.Println(i)
	}
}
