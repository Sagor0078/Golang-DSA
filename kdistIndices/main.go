import "math"

func findKDistantIndices(nums []int, key int, k int) []int {
    var result []int

    n := len(nums)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if nums[j] == key && math.Abs(float64(i) - float64(j)) <= float64(k) {
                result = append(result, i)
                break
            }
        }
    }
    return result
}