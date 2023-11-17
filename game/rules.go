package game

// CountNeighbors セルの隣接セルをカウント
func (g *Grid) CountNeighbors(x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			nx, ny := x+i, y+j
			if nx >= 0 && nx < g.rows && ny >= 0 && ny < g.cols && !(i == 0 && j == 0) {
				if g.grid[nx][ny] {
					count++
				}
			}
		}
	}
	return count
}
