func maxAbsoluteSum(nums []int) int {
    
    minpref := 0
    maxpref := 0
    sum := 0

    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        
        minpref = min(minpref, sum)
        maxpref = max(maxpref, sum)
    }
    return maxpref - minpref
}