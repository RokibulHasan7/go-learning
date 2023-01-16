func min(a int, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
func insert(intervals [][]int, newInterval []int) [][]int {
    n := len(intervals)
    Ans := [][]int{}
    i := 0
    
    for i < n && newInterval[0] > intervals[i][1] {
        curr := intervals[i]
        Ans = append(Ans, curr)
        i++
    }

    for i < n && newInterval[1] >= intervals[i][0] {
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        i++
    }
    //fmt.Println(indexOfAns)
    Ans = append(Ans, newInterval)

    for i < n {
        Ans = append(Ans, intervals[i])
        i++
    }
    return Ans
}
