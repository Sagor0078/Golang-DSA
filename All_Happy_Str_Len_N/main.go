package main

func getHappyString(n int, k int) string {
	store := []string{}
	queue := []string{"a", "b", "c"}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:] // Dequeue

		if len(curr) == n {
			store = append(store, curr)
			continue
		}

		if curr[len(curr)-1] == 'a' {
			queue = append(queue, curr+"b", curr+"c")
		} else if curr[len(curr)-1] == 'b' {
			queue = append(queue, curr+"a", curr+"c")
		} else { // 'c'
			queue = append(queue, curr+"a", curr+"b")
		}
	}

	if len(store) < k {
		return ""
	}
	return store[k-1]
}