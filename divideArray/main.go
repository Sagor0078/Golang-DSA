
import (
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	var ans [][]int

	for i := 0; i < len(nums); i += 3 {
		if i+2 < len(nums) && nums[i+2]-nums[i] > k {
			return [][]int{}
		}
		group := []int{nums[i], nums[i+1], nums[i+2]}
		ans = append(ans, group)
	}

	return ans
}
