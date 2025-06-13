
import (
	"fmt"
	"sort" 
)


func minimizeMax(nums []int, p int) int {
	n := len(nums)
	sort.Ints(nums)
	ans := 0
	lo := 0
	hi := nums[n-1] - nums[0]

	for lo <= hi {
		mid := lo + (hi-lo) / 2 

		if possible(nums, p, mid) {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func possible(nums []int, p int, diff int) bool {
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] <= diff {
			p-- 
			i++
		}
	}
	return p <= 0
}


 