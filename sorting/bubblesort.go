package sorting

import (
	"fmt"
)

// RUNTIME:
// BEST: O(N)
// AVERAGE+WORST: O(N^2)

func bubbleSort(arr []int) []int {
	swapped := true

	for swapped {
		swapped = false

		for i := 0; i < len(arr); i++ {
			if i+1 < len(arr) {
				cur := arr[i]
				next := arr[i+1]

				if cur > next {
					swapped = true
					arr[i+1] = cur
					arr[i] = next
				}
			} else {
				continue
			}
		}
	}

	return arr
}

// BubbleSortTest ...
func BubbleSortTest() {
	// Bubble Sort

	arr1 := []int{1, 2, 5, 4, 3, 6, 9, 8, 7}
	fmt.Printf("%v", bubbleSort(arr1))

	arr2 := []int{-4, -132, 48392, 93, 29, 0, 0, -10, 999}
	fmt.Printf("%v", bubbleSort(arr2))

	fmt.Println()
}
