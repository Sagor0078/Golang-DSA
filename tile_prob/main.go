package main



func dfs(f []int) int {
    ways := 0
    for i := 0; i < 26; i++ {
        if f[i] > 0 {
            f[i]--
            ways += 1 + dfs(f)
            f[i]++
        }
    }
    return ways
}
func numTilePossibilities(tiles string) int {
    f := make([]int, 26)

    for _,ch := range tiles {
        f[ch - 'A']++
    }
    
    return dfs(f)
}