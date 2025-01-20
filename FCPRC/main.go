

func firstCompleteIndex(arr []int, mat [][]int) int {
    m := len(mat)
    n := len(mat[0])

    lookup := make([][2]int, m * n + 1)

    for i:= 0; i < m; i++ {
        for j := 0; j < n; j++ {
            lookup[mat[i][j]] = [2]int {i, j}
        }
    }

    rowCount := make([]int, m)
    colCount := make([]int, n)

    for i:= 0; i < n * m; i++ {
        row, col := lookup[arr[i]][0], lookup[arr[i]][1]
        rowCount[row]++
        colCount[col]++

        if rowCount[row] == n || colCount[col] == m {
            return i
        }
    }

    return 0
}