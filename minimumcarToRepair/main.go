func canRepairAllCars(ranks []int, cars int, time int64) bool {
	var totalCars int64 = 0
	for _, r := range ranks {
		totalCars += int64(math.Sqrt(float64(time) / float64(r)))
		if totalCars >= int64(cars) {
			return true
		}
	}
	return totalCars >= int64(cars)
}

func repairCars(ranks []int, cars int) int64 {
	sort.Ints(ranks) 

	left, right := int64(1), int64(ranks[0])*int64(cars)*int64(cars)
	var answer int64 = right

	for left <= right {
		mid := left + (right-left)/2
		if canRepairAllCars(ranks, cars, mid) {
			answer = mid
			right = mid - 1 
		} else {
			left = mid + 1 
		}
	}
	return answer
}