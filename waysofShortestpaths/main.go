

import (
	"container/heap"
	"fmt"
	"math"
)

const MOD = 1e9 + 7

type Pair struct {
	dist, node int
}

type MinHeap []Pair

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist } // Min heap based on distance
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func countPaths(n int, roads [][]int) int {
	graph := make([][]Pair, n)

	for _, road := range roads {
		u, v, time := road[0], road[1], road[2]
		graph[u] = append(graph[u], Pair{time, v})
		graph[v] = append(graph[v], Pair{time, u})
	}

	return dijkstra(graph, n)
}

func dijkstra(graph [][]Pair, n int) int {
	dist := make([]int, n)
	ways := make([]int, n)

	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[0], ways[0] = 0, 1

	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, Pair{0, 0}) // {dist, node}

	for h.Len() > 0 {
		curr := heap.Pop(h).(Pair)
		d, u := curr.dist, curr.node

		if d > dist[u] {
			continue // Skip outdated distance
		}

		for _, edge := range graph[u] {
			v, time := edge.node, edge.dist
			if dist[v] > d+time {
				dist[v] = d + time
				ways[v] = ways[u]
				heap.Push(h, Pair{dist[v], v})
			} else if dist[v] == d+time {
				ways[v] = (ways[v] + ways[u]) % MOD
			}
		}
	}

	return ways[n-1]
}

