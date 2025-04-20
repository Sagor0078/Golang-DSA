import (
	"math"
)

func numRabbits(answers []int) int {
	mpp := make(map[int]int)
	total := 0

	for _, val := range answers {
		mpp[val]++
	}

	for k, v := range mpp {
		groupSize := k + 1
		total += int(math.Ceil(float64(v)/float64(groupSize))) * groupSize
	}

	return total
}
// https://leetcode.com/problems/rabbits-in-forest/?envType=daily-question&envId=2025-04-20