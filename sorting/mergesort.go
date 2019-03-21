package sorting

import "fmt"

func mergeSort(arr []int, l int, r int) []int {
	if len(arr) > 1 {
		m := (l + r) / 2
		fmt.Printf("%v", arr[:m])
		fmt.Printf("%v", arr[m:])
		left := mergeSort(arr[:m], l, m)
		right := mergeSort(arr[m:], m+1, r)

		toreturn := make([]int, len(arr))

		if len(right) == 0 {
			return left
		} else if len(left) == 0 {
			return right
		}

		i := 0
		j := 0
		for i < len(left) && j < len(right) {
			if left[i] < right[i] {
				toreturn = append(toreturn, left[i])
				i++
			} else {
				toreturn = append(toreturn, right[j])
				j++
			}
		}

		return toreturn
	}

	return arr
}

// MergeSortTest used to test quickSort method
func MergeSortTest() {

	arr0 := []int{2, 1, 3}
	mergeSort(arr0, 0, len(arr0)-1)
	fmt.Printf("%v", arr0)

	// arr1 := []int{1, 2, 5, 4, 3, 6, 9, 8, 7}
	// mergeSort(arr1, 0, len(arr1)-1)
	// fmt.Printf("%v", arr1)

	// arr2 := []int{-4, -132, 48392, 93, 29, 0, 0, -10, 999}
	// mergeSort(arr2, 0, len(arr2)-1)
	// fmt.Printf("%v", arr2)
}
