

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

func solve(word string, k int) int64 {
	left, right, cnt := 0, 0, 0
	var ans int64 = 0
	vowelMap := make(map[byte]int)

	for right < len(word) {
		if isVowel(word[right]) {
			vowelMap[word[right]]++
		} else {
			cnt++
		}

		for cnt >= k && len(vowelMap) == 5 {
			ans += int64(len(word) - right)

			if isVowel(word[left]) {
				vowelMap[word[left]]--
				if vowelMap[word[left]] == 0 {
					delete(vowelMap, word[left])
				}
			} else {
				cnt--
			}
			left++
		}
		right++
	}
	return ans
}

func countOfSubstrings(word string, k int) int64 {
	return solve(word, k) - solve(word, k+1)
}

