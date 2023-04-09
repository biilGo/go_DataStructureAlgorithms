package main

// QuickSort2三切分的快速排序
func QuickSort2(array []int, begin, end int) {
	if begin < end {
		// 三向切分函数,返回左边和右边下标
		It, gt := partition3(array, begin, end)

		// 从It到gt的部分是三切分的中间数列
		// 左边三向快排
		QuickSort2(array, begin, It-1)

		// 右边三向快排
		QuickSort2(array, gt+1, end)
	}
}

// 切分函数,并返回切分元素的下标
func partition3(array []int, begin, end int) (int, int) {
	It := begin       // 左下标从第一位开始
	gt := end         // 右下标是数组的最后一位
	i := begin + 1    // 中间下标,从第二位开始
	v := array[begin] //基准数

	// 以中间坐标为准
	for i <= gt {
		if array[i] > v {
			// 大于基准数,那么交换,右指针左移
			array[i], array[gt] = array[gt], array[i]
			gt--
		} else if array[i] < v {
			// 小于基准数,那么交换,左指针右移
			array[i], array[It] = array[It], array[i]
			It++
			i++
		} else {
			i++
		}
	}

	return It, gt
}

func main() {

}
