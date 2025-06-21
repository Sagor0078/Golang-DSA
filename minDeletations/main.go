

func minimumDeletions(word string, k int) int {
	cnt := make(map[rune]int)
	for _, ch := range word {
		cnt[ch]++
	}

	frequencies := []int{}
	for _, freq := range cnt {
		frequencies = append(frequencies, freq)
	}

	sort.Ints(frequencies)

	res := math.MaxInt32 

	potentialTargets := []int{0}
	potentialTargets = append(potentialTargets, frequencies...)
    
	for _, targetA := range potentialTargets {
		currentDeletions := 0
		for _, freqB := range frequencies {
			if freqB < targetA {
				// If frequency is less than our target 'a', delete all of them
				currentDeletions += freqB
			} else if freqB > targetA+k {
				// If frequency is greater than 'a + k', delete the excess
				currentDeletions += freqB - (targetA + k)
			}
			// If targetA <= freqB <= targetA+k, no deletions for this frequency
		}

		if currentDeletions < res {
			res = currentDeletions
		}
	}
	return res
}

