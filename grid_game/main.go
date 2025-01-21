
import (
	"fmt"
	"math"
)

func gridGame(grid [][]int) int64 {
	n := len(grid[0])
	U := int64(0)
	D := int64(0)

	for i := 1; i < n; i++ {
		U += int64(grid[0][i])
	}

	Robot2 := U

	for i := 1; i < n; i++ {
		U -= int64(grid[0][i])
		D += int64(grid[1][i-1])

		P := int64(math.Max(float64(U), float64(D)))
		if P < Robot2 {
			Robot2 = P
		}
	}

	return Robot2
}