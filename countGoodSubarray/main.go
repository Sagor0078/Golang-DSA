
func countGood(nums []int, k int) int64 {
	cnt := int64(0)
	mpp := make(map[int]int)
	left := 0

	for right := 0; right < len(nums); right++ {
		k -= mpp[nums[right]]
		mpp[nums[right]]++

		for k <= 0 {
			mpp[nums[left]]--
			k += mpp[nums[left]]
			left++
		}

		cnt += int64(left)
	}

	return cnt
}

