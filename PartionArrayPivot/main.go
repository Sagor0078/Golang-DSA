func pivotArray(nums []int, pivot int) []int {
    i, j, left, right := 0, len(nums) - 1, 0, len(nums) - 1

    res := make([]int, len(nums))

    for i:= range res{
        res[i] = pivot
    }

    for i < len(nums) {
        if nums[i] < pivot {
            res[left] = nums[i]
            left += 1
        }

        if nums[j] > pivot {
            res[right] = nums[j]
            right -= 1
        }

        i += 1
        j -= 1
    }

    return res
}