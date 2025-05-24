func findWordsContaining(words []string, x byte) []int {
    res := []int{}
    n := len(words)

    for i := 0; i < n; i++ {
        for j := 0; j < len(words[i]); j++ {
            if words[i][j] == x {
                res = append(res, i)
                break
            }
        }
    }
    return res
}