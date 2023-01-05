func findMinArrowShots(points [][]int) int {
    // Sort 2D slice according to the 2nd element
    sort.Slice(points, func(i, j int) bool {
        return points[i][1] < points[j][1]
    })

    ans := 1

    endPoint := points[0][1]
    for i := 1; i < len(points); i++ {
        if endPoint < points[i][0] {
            ans++
            endPoint = points[i][1]
        }
    }
    return ans
}
