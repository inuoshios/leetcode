package main

func firstBadVersion(n int) int {
	for i := 0; i < n; i++ {
		if isBadVersion(i) {
			return i
		}
	}
	return n
}
