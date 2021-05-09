//Package sorts a package for demonstrating sorting algorithms in Go
// 插入排序
// https://zh.wikipedia.org/zh/%E5%B8%8C%E5%B0%94%E6%8E%92%E5%BA%8F
// https://en.wikipedia.org/wiki/Shellsort
// https://baike.baidu.com/item/%E5%B8%8C%E5%B0%94%E6%8E%92%E5%BA%8F

package sorts

func ShellSort(arr []int) []int {
	for d := int(len(arr) / 2); d > 0; d /= 2 {
		for i := d; i < len(arr); i++ {
			for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
				arr[j], arr[j-d] = arr[j-d], arr[j]
			}
		}
	}
	return arr
}
