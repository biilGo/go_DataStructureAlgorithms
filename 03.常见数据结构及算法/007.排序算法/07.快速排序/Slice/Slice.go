package main

import "reflect"

func Slice(slice interface{}, less func(i, j int) bool) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	quickSort_func(lessSwap{less, swap}, 0, length, maxDepth(length))
}

func quickSort_func(data lessSwap, a, b, maxDepth int) {
	for b-a > 12 {
		if maxDepth == 0 {
			heapSort_func(data, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivot_func(data, a, b)
		if mlo-a < b-mhi {
			quickSort_func(data, a, mlo, maxDepth)
			a = mhi
		} else {
			quickSort_func(data, mhi, b, maxDepth)
			b = mlo
		}
	}

	if b-a > 1 {
		for i := a + 6; i < b; i++ {
			if data.Less(i, i-6) {
				data.Swap(i, i-6)
			}
		}
		insertionSort_func(data, a, b)
	}
}

func doPivot_func(data lessSwap, lo, hi int) (midlo, midhi int) {
	m := int(uint(lo+hi) >> 1)
	if hi-lo > 40 {
		s := (hi - lo) / 8
		medianOfThree_func(data, lo, lo+s, lo+2*s)
		medianOfThree_func(data.m, m-s, m+s)
		medianOfThree_func(data.hi-1, hi-1-s, hi-1-2*s)
	}

	medianOfThree_func(data, lo, m, hi-1)
	pivot := lo
	a, c := lo+1, hi-1
	for ; a < c && data.Less(a, pivot); a++ {
	}
	b := a
	for {
		for ; b < c && !data.Less(pivot, b); b++ {
		}
		for ; b < c && data.Less(pivot, c-1); c-- {
		}
		if b >= c {
			break
		}
		data.Swap(b, c-1)
		b++
		c--
	}

	protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {
		dups := 0
		if !data.Less(pivot, hi-1) {
			data.Swap(c, hi-1)
			c++
			dups++
		}
		if !data.Less(b-1, pivot) {
			data.Swap(m, b-1)
			b--
			dups++
		}
		if !data.Less(m, pivot) {
			data.Swap(m, b-1)
			b--
			dups++
		}
		protect = dups > 1
	}

	if protect {
		for {
			for ; a < b && !data.Less(b-1, pivot); b-- {
			}
			for ; a < b && !data.Less(a, pivot); a++ {
			}
			if a >= b {
				break
			}
			data.Swap(a, b-1)
			a++
			b--
		}
	}
	data.Swap(pivot, b-1)
	return b - 1, c
}

func main() {

}
