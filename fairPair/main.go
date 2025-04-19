
import (
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	var answer int64 = 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		low := lowerBound(nums, i+1, n, lower-nums[i])
		up := upperBound(nums, i+1, n, upper-nums[i])
		answer += int64(up - low)
	}
	return answer 
}

func lowerBound(nums []int, l, r, target int) int {
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func upperBound(nums []int, l, r, target int) int {
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
