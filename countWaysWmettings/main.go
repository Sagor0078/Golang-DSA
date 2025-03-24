import (
	"sort"
)

func countDays(days int, meetings [][]int) int {
	if len(meetings) == 0 {
		return days
	}

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	busyDays := 0
	prevStart, prevEnd := -1, -1

	for _, meeting := range meetings {
		start, end := meeting[0], meeting[1]
		if start <= prevEnd {
			prevEnd = max(prevEnd, end)
		} else {
			busyDays += (prevEnd - prevStart + 1)
			prevStart, prevEnd = start, end
		}
	}
	
	busyDays += (prevEnd - prevStart + 1)
	
	return (days - busyDays) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
