

func dfs(node int, vis []bool, adj [][]int, nodes *int, edges *int) {
    vis[node] = true
    *nodes++
    *edges += len(adj[node])

    for _, neigh := range adj[node] {
        if !vis[neigh] {
            dfs(neigh, vis, adj, nodes, edges)
        }
    }
}

func countCompleteComponents(n int, edges [][]int) int {
    adj := make([][]int, n)
    for _, e := range edges {
        adj[e[0]] = append(adj[e[0]], e[1])
        adj[e[1]] = append(adj[e[1]], e[0])
    }

    vis := make([]bool, n)
    components := 0

    for i := 0; i < n; i++ {
        if !vis[i] {
            nodes, edges := 0, 0
            dfs(i, vis, adj, &nodes, &edges)
            if edges == nodes*(nodes-1) {
                components++
            }
        }
    }
    return components
}

