package sorting

// func QuickSort(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	pivot := arr[len(arr)/2]
// 	var left, right []int

// 	// Partitioning the slice into left and right sub-slices
// 	for _, val := range arr {
// 		if val < pivot {
// 			left = append(left, val)
// 		} else if val > pivot {
// 			right = append(right, val)
// 		}
// 	}

// 	// Recursively sort the left and right sub-slices
// 	left = QuickSort(left)
// 	right = QuickSort(right)

// 	// Combine the results with the pivot
// 	return append(append(left, pivot), right...)
// }

func QuickSort(arr []int) {
	stack := []struct{ low, high int }{{0, len(arr) - 1}}

	for len(stack) > 0 {
		// Pop a range from the stack
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		low, high := top.low, top.high
		if low >= high {
			continue
		}

		// Partition the array
		pivot := partition(arr, low, high)

		// Push left and right subarrays onto the stack
		stack = append(stack, struct{ low, high int }{low, pivot - 1})
		stack = append(stack, struct{ low, high int }{pivot + 1, high})
	}
}

// partition rearranges elements based on the pivot and returns its position
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
