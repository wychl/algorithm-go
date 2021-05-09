//Package sorts a package for demonstrating sorting algorithms in Go
// https://zh.wikipedia.org/wiki/%E9%80%89%E6%8B%A9%E6%8E%92%E5%BA%8F

package sorts

// 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
// 再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾
func SelectionSort(arr []int) []int {

	// i为已排序序列的最后一个
	for i := 0; i < len(arr); i++ {
		min := i

		// 从剩余未排序元素中寻找小于i的元素
		// 如果找到交换两个元素
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		tmp := arr[i]
		arr[i] = arr[min]
		arr[min] = tmp
	}
	return arr
}
