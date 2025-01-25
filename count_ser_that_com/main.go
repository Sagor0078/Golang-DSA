package main

func countServers(grid [][]int) int {
	N := len(grid)
	M := len(grid[0])

	row := make([]int, N)
	col := make([]int, M)

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if grid[i][j] == 1 {
				row[i]++
				col[j]++
			}
		}
	}
	cnt := 0

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if grid[i][j] == 1 {
				if row[i] > 1 || col[j] > 1 {
					cnt++
				}
			}
		}
	}

	return cnt
}
