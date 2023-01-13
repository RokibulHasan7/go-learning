func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}
func dfs(currNode int, adj [][]int, s string, Ans *int) int {
    mxLen, secondMxLen := 0, 0
    for i := 0; i < len(adj[currNode]); i++ {
        child := adj[currNode][i]
        mxLenFromChild := dfs(child, adj, s, Ans)
        if s[child] == s[currNode] {
            continue
        }

        //mxLenFromChild := dfs(child, adj, s, Ans)
        if mxLenFromChild > mxLen {
            secondMxLen = mxLen
            mxLen = mxLenFromChild
        } else if mxLenFromChild > secondMxLen {
            secondMxLen = mxLenFromChild
        }
    }
    *Ans = max(*Ans, mxLen + secondMxLen + 1)
    return mxLen + 1
}
func longestPath(parent []int, s string) int {
    sz := len(parent)
    adj := make([][]int, sz)
    for i := 1; i < sz; i++ {
        u := parent[i]
        v := i
        adj[u] = append(adj[u], v)
    }

    Ans := 1
    dfs(0, adj, s, &Ans)

    return Ans
}
