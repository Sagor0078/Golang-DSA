func longestNiceSubarray(nums []int) int {
    left, right, res, bits := 0, 0, 0, 0
    
    for right < len(nums) {

        for bits & nums[right] != 0 {
            bits ^= nums[left]
            left++
        }

        bits |= nums[right]
        res = max(res, right - left + 1)
        right++
    }
    return res
}