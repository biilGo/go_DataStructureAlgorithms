package main

import "fmt"

func main() {
	// 新建一个容量为4的字典map
	m := make(map[string]int64, 4)

	// 放3个键值对
	m["dog"] = 1
	m["cat"] = 2
	m["hen"] = 3

	fmt.Println(m)

	// 查找hen
	which := "hen"
	v, ok := m[which]
	if ok {
		// 找到了
		fmt.Println("find:", which, "value:", v)
	} else {
		// 找不到
		fmt.Println("not find:", which)
	}

	// 查找ccc
	which2 := "ccc"
	v2, ok := m[which2]
	if ok {
		// 找到了
		fmt.Println("find:", which2, "value", v2)
	} else {
		// 找不到
		fmt.Println("not find:", which2)
	}

}
