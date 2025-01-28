package main

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

func dfs(r, c, M, N int, grid [][]int) int {
	if r < 0 || r >= M || c < 0 || c >= N || grid[r][c] == 0 {
		return 0
	}

	ans := grid[r][c]
	grid[r][c] = 0

	for i := 0; i < 4; i++ {
		rr := r + dx[i]
		cc := c + dy[i]
		ans += dfs(rr, cc, M, N, grid)
	}
	return ans
}

func findMaxFish(grid [][]int) int {
	M := len(grid)
	if M == 0 {
		return 0 // Handle empty grid case
	}
	N := len(grid[0])

	ans := 0
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] > 0 {
				ans = max(ans, dfs(i, j, M, N, grid))
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
