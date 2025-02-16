package main

func constructDistancedSequence(n int) []int {
    res := make([]int, 2 * n -1)
    used := make([]bool, n + 1)

    var backtrack func(i int) bool 

    backtrack = func(i int) bool {
        if i == len(res) {
            return true
        }

        if res[i] != 0 {
            return backtrack(i + 1)
        }

        for x := n; x >= 1; x-- {
            if used[x] {
                continue
            }

            if x > 1 && (i + x >= len(res) || res[i + x] != 0) {
                continue
            }

            used[x] = true
            res[i] = x
            if x > 1 {
                res[i + x] = x
            }

            if backtrack(i + 1) {
                return true
            }

            used[x] = false
            res[i] = 0
            if x > 1 {
                res[i + x] = 0
            }
        }
        return false
    }

    backtrack(0)
    return res
}