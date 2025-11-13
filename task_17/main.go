package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func binarySearchRecursive(arr []int, target int, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return binarySearchRecursive(arr, target, mid+1, right)
	} else {
		return binarySearchRecursive(arr, target, left, mid-1)
	}
}

func main() {
	nums := []int{1, 3, 5, 7, 9, 11}
	fmt.Println(binarySearch(nums, 7))                          // → 3
	fmt.Println(binarySearch(nums, 4))                          // → -1
	fmt.Println(binarySearchRecursive(nums, 9, 0, len(nums)-1)) // → 4
	fmt.Println(binarySearchRecursive(nums, 2, 0, len(nums)-1)) // → -1
}
