package main

func countBadPairs(nums []int) int64 {
	n := int64(len(nums))
	good_pairs := int64(0)
	hash := make(map[int]int64)

	for i, num := range nums {
		key := num - i

		if count, exits := hash[key]; exits {
			good_pairs += count
		}
		hash[key]++
	}

	total_pairs := n * (n - 1) / 2
	return total_pairs - good_pairs
}
