func findMissingAndRepeatedValues(grid [][]int) []int {
    mp := make(map[int]int)
    sum := 0
    n := len(grid)
    duplicate := 0

    // Count occurrences and find duplicate
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            mp[grid[i][j]]++
            sum += grid[i][j]
            if mp[grid[i][j]] > 1 {
                duplicate = grid[i][j]
            }
        }
    }

    total := n * n
    missing := total * (total + 1) / 2 - (sum - duplicate)

    return []int{duplicate, missing}
}