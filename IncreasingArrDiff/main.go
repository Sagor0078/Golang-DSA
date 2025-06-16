func maximumDifference(nums []int) int {
    max_element := -1
    curr_left := nums[0]
    for i := 0; i < len(nums); i++ {
       if nums[i] > curr_left {
            max_element = max(max_element, nums[i] - curr_left)
       }
       curr_left = min(curr_left, nums[i])
    }
    return max_element
}