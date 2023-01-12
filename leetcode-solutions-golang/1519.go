func dfs(currNode int, parentNode int, adj [][]int, labels string, Ans []int) []int {
    parentLabel := make([]int, 26)
    parentLabel[labels[currNode] - 'a'] = 1;
    for _, child := range adj[currNode] {
        if child == parentNode {
            continue
        }
        childLabel := make([]int, 26)
        childLabel = dfs(child, currNode, adj, labels, Ans)
        for i := 0; i < 26; i++ {
            parentLabel[i] += childLabel[i]
        }
    }

    Ans[currNode] = parentLabel[labels[currNode] - 'a']
    return parentLabel
}
func countSubTrees(n int, edges [][]int, labels string) []int {
    Ans := make([]int, n)
    adj := make([][]int, n)
    for _, edge := range edges {
        u := edge[0]
        v := edge[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }
    dfs(0, -1, adj, labels, Ans)
    return Ans
}
