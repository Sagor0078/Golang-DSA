package main

import (
	"container/heap"
)

type Cell struct {
	height int
	row    int
	col    int
}

type MinHeap []Cell

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].height < h[j].height
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Cell))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func highestPeak(isWater [][]int) [][]int {
	M := len(isWater)
	N := len(isWater[0])

	res := make([][]int, M)
	for i := range res {
		res[i] = make([]int, N)
		for j := range res[i] {
			res[i][j] = -1
		}
	}

	h := &MinHeap{}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if isWater[i][j] == 1 {
				res[i][j] = 0
				heap.Push(h, Cell{height: 0, row: i, col: j})
			}
		}
	}

	for h.Len() > 0 {
		cell := heap.Pop(h).(Cell)
		height := cell.height
		r := cell.row
		c := cell.col

		for i := 0; i < 4; i++ {
			nr, nc := r+dx[i], c+dy[i]

			if nr >= 0 && nr < M && nc >= 0 && nc < N && res[nr][nc] == -1 {
				res[nr][nc] = height + 1
				heap.Push(h, Cell{height: height + 1, row: nr, col: nc})
			}
		}
	}

	return res
}
