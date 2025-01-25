package main

func eventualSafeNodes(graph [][]int) []int {
	colors := make([]int, len(graph))
	res := make([]int, 0)
	for i, _ := range graph {
		if dfs(i, colors, graph) {
			res = append(res, i)
		}
	}
	return res
}

func dfs(curr int, colors []int, graph [][]int) bool {
	if colors[curr] > 0 {
		return colors[curr] == 2
	}
	colors[curr] = 1
	for _, adj := range graph[curr] {
		if !dfs(adj, colors, graph) {
			return false
		}
	}
	colors[curr] = 2
	return true
}
