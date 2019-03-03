package sorting

import "fmt"

func quickSort(arr []int, low int, high int) {

	if low < high {
		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	// this will take the last element as the pivot

	// element to be placed at the right most position
	pivot := arr[high]

	// index of smaller element
	i := low - 1

	for j := low; j <= high-1; j++ {
		// If current element is smaller than or
		// equal to pivot
		if arr[i] <= pivot {
			i++
			swapInArray(arr, arr[i], arr[j])
		}
	}

	swapInArray(arr, arr[i+1], arr[high])
	return (i + 1)

}

func swapInArray(arr []int, first int, second int) {
	temp := arr[first]
	arr[first] = arr[second]
	arr[second] = temp
}

// QuickSortTest used to test quickSort method
func QuickSortTest() {

	arr1 := []int{1, 2, 5, 4, 3, 6, 9, 8, 7}
	quickSort(arr1, 0, len(arr1)-1)
	fmt.Printf("%v", arr1)

	arr2 := []int{-4, -132, 48392, 93, 29, 0, 0, -10, 999}
	quickSort(arr2, 0, len(arr2)-1)
	fmt.Printf("%v", arr2)
}
