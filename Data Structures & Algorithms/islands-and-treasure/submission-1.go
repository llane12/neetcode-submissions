// Multi-source BFS
func islandsAndTreasure(grid [][]int) {
    q := make([][3]int, 0)
    visited := make(map[[2]int]struct{})

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 0 {
				q = append(q, [3]int{r, c, 0})
				// BFS - mark visited WHEN ADDING TO QUEUE!
        		visited[[2]int{r, c}] = struct{}{}
			}
		}
	}

    for len(q) > 0 {
		// Pop cell from queue
        cell := q[0]
        q = q[1:]
		
		r := cell[0]
		c := cell[1]
		dist := cell[2]
		
		// Update the distance from starting cell(s)
		grid[r][c] = dist
        
		// Find neighbour cells
        neighbours := [4][2]int{
            {r + 1, c},
            {r - 1, c},
            {r, c + 1},
            {r, c - 1},
        }

        for _, nei := range neighbours {
            if _, seen := visited[nei]; !seen && canVisit(nei[0], nei[1], grid) {
                q = append(q, [3]int{nei[0], nei[1], dist + 1})
				// BFS - mark visited WHEN ADDING TO QUEUE!
        		visited[[2]int{nei[0], nei[1]}] = struct{}{}
            }
        }
    }
}

func canVisit(r, c int, grid [][]int) bool {
    return r >= 0 && c >= 0 && r < len(grid) && c < len(grid[0]) && grid[r][c] > -1
}
