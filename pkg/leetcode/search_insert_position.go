package leetcode

// Search Insert Position
// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
// You must write an algorithm with O(log n) runtime complexity.

func SearchInsertPosition(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		var (
			midIdx = (left + right) / 2
			midNum = nums[midIdx]
		)

		if midNum == target {
			return midIdx
		}

		if midNum < target {
			left = midIdx + 1
		} else if target < midNum {
			right = midIdx - 1
		}
	}

	return left
}
