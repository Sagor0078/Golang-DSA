

func canPartition(nums []int) bool {
    total := 0
    for _, num := range nums {
        total += num
    }

    if total%2 != 0 {
        return false
    }

    target := total / 2
    n := len(nums)

    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, target+1)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }

    var dfs func(i, sum int) bool
    dfs = func(i, sum int) bool {
        if sum == target {
            return true
        }
        if i >= n || sum > target {
            return false
        }
        if dp[i][sum] != -1 {
            return dp[i][sum] == 1
        }

        take := dfs(i+1, sum+nums[i])
        skip := dfs(i+1, sum)

        if take || skip {
            dp[i][sum] = 1
            return true
        } else {
            dp[i][sum] = 0
            return false
        }
    }

    return dfs(0, 0)
}
