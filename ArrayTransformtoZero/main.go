package main

import (
	"container/heap"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxIntHeap []int

func (h MaxIntHeap) Len() int           { return len(h) }
func (h MaxIntHeap) Less(i, j int) bool { return h[i] > h[j] } // Max-heap
func (h MaxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxRemoval(nums []int, queries [][]int) int {
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][0] < queries[j][0]
	})

	available := &MaxIntHeap{}
	heap.Init(available)

	assigned := &IntHeap{}
	heap.Init(assigned)

	count := 0
	k := 0 // Pointer for queries

	for time := 0; time < len(nums); time++ {
		for assigned.Len() > 0 && (*assigned)[0] < time {
			heap.Pop(assigned)
		}

		for k < len(queries) && queries[k][0] <= time {
			heap.Push(available, queries[k][1])
			k++
		}

		for assigned.Len() < nums[time] && available.Len() > 0 && (*available)[0] >= time {
			heap.Push(assigned, heap.Pop(available).(int))
			count++
		}

		if assigned.Len() < nums[time] {
			return -1
		}
	}

	return len(queries) - count
}