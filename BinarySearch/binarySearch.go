package main

import (
	"sort"
)

func binarySearch(nums []int, target int) int {
	sort.Ints(nums)
	firstIndex := 0
	lastIndex := len(nums) - 1

	for i, num := range nums {
		midIndex := (firstIndex + lastIndex) / 2

		if num == target {
			return i
		} else if num < target {
			firstIndex = midIndex + 1
		} else {
			lastIndex = midIndex - 1
		}
	}
	return -1
}
