type Cell struct {
	height int
	row    int
	col    int
}

// PriorityQueue implements heap.Interface
type PriorityQueue []Cell

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].height < pq[j].height }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Cell)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}

	m, n := len(heightMap), len(heightMap[0])
	pq := &PriorityQueue{}
	heap.Init(pq)
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				heap.Push(pq, Cell{height: heightMap[i][j], row: i, col: j})
				visited[i][j] = true
			}
		}
	}

	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	waterTrapped := 0
	waterLevel := math.MinInt

	for pq.Len() > 0 {
		cell := heap.Pop(pq).(Cell)
		waterLevel = max(waterLevel, cell.height)

		for _, d := range directions {
			newRow, newCol := cell.row+d[0], cell.col+d[1]
			if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && !visited[newRow][newCol] {
				visited[newRow][newCol] = true
				neighborHeight := heightMap[newRow][newCol]

				if neighborHeight < waterLevel {
					waterTrapped += waterLevel - neighborHeight
				}

				heap.Push(pq, Cell{height: neighborHeight, row: newRow, col: newCol})
			}
		}
	}

	return waterTrapped
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}