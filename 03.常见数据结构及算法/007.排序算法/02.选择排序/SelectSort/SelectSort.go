package main

import "fmt"

func SelectSort(list []int) {
	n := len(list)
	// 进行N-1轮迭代
	for i := 0; i < n-1; i++ {
		// 每次从第i位开始,找到最小的元素
		min := list[i] // 最小数
		minIndex := i  // 最小数的小标
		for j := i + 1; j < n; j++ {
			if list[j] < min {
				// 如果找到的数比上次的还小,那么最小的数变为它
				min = list[i]
				minIndex = j
			}
		}

		// 这一轮找到的最小数的下标不等于最开始的下标,交换元素
		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	SelectSort(list)
	fmt.Println(list)
}

// 每一轮迭代,我们都会维持这一轮最小数:min和最小数的下标:minIndex,然后开始扫描,如果扫描的数比该数小,那么替换掉最小数和最小数小标,扫描完后判断是否应该交换,然后交换:list[i], list[minIndex] = list[minIndex], list[i]
