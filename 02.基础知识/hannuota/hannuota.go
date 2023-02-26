package main

import "fmt"

var total = 0

// 汉诺塔
// 一开始A杆上有N个盘子,B和C杆都没有盘子

func towner(n int, a, b, c string) {
	if n == 1 {
		total = total + 1
		fmt.Println(a, "->", c)
		return
	}

	towner(n-1, a, c, b)
	total = total + 1
	fmt.Println(a, "->", c)
	towner(n-1, b, a, c)
}

func main() {
	n := 4 // 64个盘子
	a := "a"
	b := "b"
	c := "c"
	towner(n, a, b, c)

	// 当n=1时,移动次数为1
	// 当n=2时,移动次数为3
	// 当n=3时,移动次数为7
	// 当n=4时,移动次数为15
	fmt.Println(total)
}
