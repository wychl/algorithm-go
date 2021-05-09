// Implementation of basic bubble sort algorithm
// Reference: https://en.wikipedia.org/wiki/Bubble_sort

package sorts

func bubbleSort(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
	return arr
}

// The bubble sort algorithm can be optimized by observing that the n-th pass finds the n-th largest element and puts it into its final place. So, the inner loop can avoid looking at the last n − 1 items when running for the n-th time:
func bubbleSort2(arr []int) []int {
	swapped := true
	n := len(arr)
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
		n = n - 1
	}
	return arr
}

// More generally, it can happen that more than one element is placed in their final position on a single pass. In particular, after every pass, all elements after the last swap are sorted, and do not need to be checked again. This allows to skip over many elements, resulting in about a worst case 50% improvement in comparison count (though no improvement in swap counts), and adds very little complexity because the new code subsumes the "swapped" variable:
func bubbleSort3(arr []int) []int {
	swapped := true
	n := len(arr)
	for swapped {
		newn := 0
		swapped = false
		for i := 0; i < n-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
				newn = i + 1
			}
		}
		n = newn
	}
	return arr
}
