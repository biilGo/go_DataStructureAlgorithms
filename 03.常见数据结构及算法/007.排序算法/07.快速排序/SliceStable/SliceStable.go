package main

import "reflect"

// 其中 SliceStable 是稳定排序，使用了插入排序和归并排序：
func SliceStable(slice interface{}, less func(i, j int) bool) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	stable_func(lessSwap{less, swap}, rv.Len())
}

func stable_func(data lessSwap, n int) {
	blockSize := 20
	a, b := 0, blockSize
	for b <= n {
		insertionSort_func(data, a, b)
		a = b
		b += blockSize
	}
	insertionSort_func(data, a, n)
	for blockSize < n {
		a, b = 0, 2*blockSize
		for b <= n {
			symMerge_func(data, a, a+blockSize, b)
			a = b
			b += 2 * blockSize
		}
		if m := a + blockSize; m < n {
			symMerge_func(data, a, m, n)
		}
		blockSize *= 2
	}
}

func main() {

}
