

type UnionFind struct {
    parent []int
    andVal []int
}

func NewUnionFind(n int) *UnionFind {
    parent := make([]int, n)
    andVal := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
        andVal[i] = -1 // Initialize AND values
    }
    return &UnionFind{parent, andVal}
}

func (uf *UnionFind) Find(x int) int {
    if uf.parent[x] != x {
        uf.parent[x] = uf.Find(uf.parent[x])
    }
    return uf.parent[x]
}

func (uf *UnionFind) Union(x, y, w int) {
    rootX := uf.Find(x)
    rootY := uf.Find(y)
    if rootX != rootY {
        uf.parent[rootY] = rootX
        uf.andVal[rootX] &= uf.andVal[rootY]
    }
    uf.andVal[rootX] &= w
}

func minimumCost(n int, edges [][]int, queries [][]int) []int {
    uf := NewUnionFind(n)

    for _, edge := range edges {
        u, v, w := edge[0], edge[1], edge[2]
        uf.Union(u, v, w)
    }

    result := make([]int, len(queries))
    for i, q := range queries {
        u, v := q[0], q[1]
        rootU := uf.Find(u)
        rootV := uf.Find(v)
        if u == v {
            result[i] = 0
        } else if rootU != rootV {
            result[i] = -1
        } else {
            result[i] = uf.andVal[rootU]
        }
    }
    return result
}

