package main

func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
	adj := make(map[int][]int)

	for _, edge := range prerequisites {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
	}

	ans := make([]bool, len(queries))

	for i, q := range queries {
		src := q[0]
		target := q[1]

		queue := []int{src}
		vis := make([]bool, n)

		vis[src] = true
		found := false
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			if curr == target {
				found = true
				break
			}

			for _, neigh := range adj[curr] {
				if !vis[neigh] {
					vis[neigh] = true
					queue = append(queue, neigh)
				}
			}
		}
		ans[i] = found
	}
	return ans
}
