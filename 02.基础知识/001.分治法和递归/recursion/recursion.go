package main

import "fmt"

func Rescuvie(n int) int {
	if n == 0 {
		return 1
	}

	return n * Rescuvie(n-1)
}

func main() {
	fmt.Println(Rescuvie(5))
}
