
func pacificAtlantic(heights [][]int) [][]int {
    if len(heights) == 0 || len(heights[0]) == 0 {
        return [][]int{}
    }

    rows, cols := len(heights), len(heights[0])

    ocean := make([][]int, rows)
    for i := range ocean {
        ocean[i] = make([]int, cols)
    }

    directions := [][]int{
        {0, 1}, {0, -1}, {1, 0}, {-1, 0},
    }

    oceanlist := [][]int{}

    // Pacific: top and left
    // Atlantic: bottom and right
    for y := 0; y < cols; y++ {
        ocean[0][y] |= 1
        oceanlist = append(oceanlist, []int{0, y})

        ocean[rows-1][y] |= 2
        oceanlist = append(oceanlist, []int{rows-1, y})
    }

    for x := 0; x < rows; x++ {
        ocean[x][0] |= 1
        oceanlist = append(oceanlist, []int{x, 0})

        ocean[x][cols-1] |= 2
        oceanlist = append(oceanlist, []int{x, cols-1})
    }

    for len(oceanlist) > 0 {
        pos := oceanlist[0]
        oceanlist = oceanlist[1:]

        for _, d := range directions {
            x := pos[0] + d[0]
            y := pos[1] + d[1]

            if x < 0 || x >= rows ||
                y < 0 || y >= cols {
                continue
            }

            if heights[x][y] < heights[pos[0]][pos[1]] {
                continue
            }

            old := ocean[x][y]
            newState := old | ocean[pos[0]][pos[1]]

            if newState != old {
                ocean[x][y] = newState
                oceanlist = append(oceanlist, []int{x, y})
            }
        }
    }

    flow := [][]int{}

    for x := 0; x < rows; x++ {
        for y := 0; y < cols; y++ {
            if ocean[x][y] == 3 {
                flow = append(flow, []int{x, y})
            }
        }
    }

    return flow
}
