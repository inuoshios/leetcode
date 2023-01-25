package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Binary Search")
	search := []int{3, 1, 2, 6, 5, 7, 4, 3, 9, 7, 8, 10}
	searchData := binarySearch(search, 10)
	fmt.Println(searchData)
}

func binarySearch(nums []int, target int) bool {
	sort.Ints(nums)
	low := 0
	high := len(nums) - 1

	for _, num := range nums {
		midIndex := (high + low) / 2
		if num == target {
			return true
		} else if num > nums[midIndex] {
			high = midIndex + 1
		} else {
			low = midIndex - 1
		}
	}
	return false
}
