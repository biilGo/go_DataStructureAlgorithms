package main

import "fmt"

func RescuvieTail(n int, a int) int {
	if n == 1 {
		return a
	}

	return RescuvieTail(n-1, a*n)
}

func main() {
	fmt.Println(RescuvieTail(5, 1))
}
