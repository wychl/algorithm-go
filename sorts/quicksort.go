//Package sorts a package for demonstrating sorting algorithms in Go
// https://en.wikipedia.org/wiki/Quicksort
// https://zh.wikipedia.org/zh/%E5%BF%AB%E9%80%9F%E6%8E%92%E5%BA%8F

package sorts

func QuickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	q := partition(array)
	QuickSort(array[:q])
	QuickSort(array[q+1:])
	return array
}

// 序列分割-原地分割算法
func partition(array []int) int {
	r := len(array) - 1
	x := array[r] // 使用最后一个值作为基准值
	i := -1
	for j := 0; j < r; j++ {
		if array[j] <= x {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[r] = array[r], array[i+1]
	return i + 1
}
