// InsertSort改进:当数组规模小时使用直接插入排序
func InsertSort(list []int) {
	n := len(list)

	// 进行N-1轮迭代
	for i := 1; i <= n-1; i++ {
		deal := list[i] //待排序的数
		j := i - 1      // 待排序的数左边的第一个数的位置

		// 如果第一次比较,比左边的已排好序的第一个数小,那么进入处理
		if deal < list[j] {
			// 一直往左边找,比待排序大的数都往后挪,腾空位给待排序插入
			for ; j >= 0 && deal < list[j]; j-- {
				list[j+1] = list[j] // 某数后移,给待排序留空位
			}
			list[j+1] = deal // 结束了,待排序的数插入空位
		}
	}
}

func QuickSort1(array []int, begin, end int) {
	if begin < end {
		// 当数组小于4时使用直接插入排序
		if end-begin <= 4 {
			InsertSort(array[begin : end+1])
			return
		}

		// 进行切分
		loc := partition(array, begin, end)

		// 对左部分进行快排
		QuickSort1(array, begin, loc-1)

		// 对右部分进行快排
		QuickSort1(array, loc+1, end)
	}
}