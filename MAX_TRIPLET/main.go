

func maximumTripletValue(nums []int) int64 {
    n := len(nums)
    maxNum, maxDiff, ans := int64(nums[0]), int64(0), int64(0)
    
    for i := 1; i < n; i++ {
        num := int64(nums[i])
        
        ans = max(ans, maxDiff*num)
        
        maxDiff = max(maxDiff, maxNum-num)
        
        maxNum = max(maxNum, num)
    }
    
    return ans
}

func max(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}