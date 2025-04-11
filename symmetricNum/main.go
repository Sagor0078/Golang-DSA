

func countSymmetricIntegers(low int, high int) int {
	count := 0
	for x := low; x <= high; x++ {
		num := fmt.Sprintf("%d", x)
		length := len(num)

		if length%2 != 0 {
			continue
		}

		half := length / 2
		firstHalf, secondHalf := 0, 0

		for i := 0; i < half; i++ {
			firstHalf += int(num[i] - '0')
		}

		for i := half; i < length; i++ {
			secondHalf += int(num[i] - '0')
		}

		if firstHalf == secondHalf {
			count++
		}
	}
	return count
}
