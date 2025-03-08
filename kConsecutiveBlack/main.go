

func minimumRecolors(blocks string, k int) int {
    n := len(blocks)
    left, right, changes := 0, 0, 0
    minOperation := math.MaxInt32
    for right < n {
        if blocks[right] == 'W'{
            changes++;
        }

        if right - left + 1 == k {
            minOperation = min(minOperation, changes)

            if blocks[left] == 'W' {
                changes--;
            }
            left++;
        }
        right++;
    }
    return minOperation
}