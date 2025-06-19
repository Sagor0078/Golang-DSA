func partitionArray(nums []int, k int) int {
    sort.Ints(nums)
    rec := nums[0]
    ans := 1
    for _, num := range nums {
        if num - rec > k {
            ans++
            rec = num
        }
    }
    return ans
}