// 冒泡排序

package main

import "fmt"

func BubbleSort(list []int) {
	n := len(list)

	// 进行N-1轮迭代
	for i := n - 1; i > 0; i-- {

		// 引入didSwap变量,如果在一轮中该变量值没有变换,那么表示数列是有序的,所以不需要交换,也就是说在最好的情况下:
		// 对已经排好序的数列进行冒泡排序,只需比较`N`次,最好时间复杂度从O(n^2) 骤减为 O(n)

		didSwap := false

		// 每次从第一位开始比较,比较到第i位就不比较了,因为前一轮该位已经有序了
		for j := 0; j < i; j++ {
			// 如果前面的数比后面的大,那么交换
			if list[j] > list[j+1] {

				/*
					很多编程语言不允许使用list[j], list[j+1] = list[j+1], list[j]来交换两个值,交换两个值时必须建一个临时变量a来作为一个过渡:
					a := list[j+1]
					list[j+1] = list[j]
					list[j] = a
				*/

				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}
		}

		// 如果在一轮中没有交换过,那么已经排序好了,直接返回
		if !didSwap {
			return
		}
	}
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	BubbleSort(list)

	// 由于切片list会原地排序,排序函数不需要返回任何值,处理完后可以直接打印:
	fmt.Println(list)
}
