package main

import "fmt"

func find(parent []int, i int) int {
	if parent[i] != i {
		parent[i] = find(parent, parent[i])
	}
	return parent[i]
}

func join(parent, rank []int, u, v int) {
	rootU := find(parent, u)
	rootV := find(parent, v)

	if rootU != rootV {
		if rank[rootU] > rank[rootV] {
			parent[rootV] = rootU
		} else if rank[rootU] < rank[rootV] {
			parent[rootU] = rootV
		} else {
			parent[rootV] = rootU
			rank[rootU]++
		}
	}
}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	rank := make([]int, n+1)

	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if find(parent, u) == find(parent, v) {
			return edge
		}
		join(parent, rank, u, v)
	}
	return nil
}

func main() {
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	}
	result := findRedundantConnection(edges)
	fmt.Println(result) // Output: [2, 3]
}
