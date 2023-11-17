package game

import (
	"errors"
	"log"
)

type Grid struct {
	rows, cols int
	grid       [][]bool
}

// NewGrid 新しいグリッドを作成
func NewGrid(rows, cols int) (*Grid, error) {
	if rows <= 0 || cols <= 0 {
		return nil, errors.New("invalid grid size")
	}

	grid := make([][]bool, rows)
	for i := range grid {
		grid[i] = make([]bool, cols)
	}
	return &Grid{rows: rows, cols: cols, grid: grid}, nil
}

// SetCell セルの状態を設定
func (g *Grid) SetCell(x, y int, state bool) {
	if x >= 0 && x < g.rows && y >= 0 && y < g.cols {
		g.grid[x][y] = state
	}
}

// GetCell セルの状態を取得
func (g *Grid) GetCell(x, y int) bool {
	if x >= 0 && x < g.rows && y >= 0 && y < g.cols {
		return g.grid[x][y]
	}
	return false
}

// NextGeneration 次世代のグリッドを計算
func (g *Grid) NextGeneration() *Grid {
	nextGen, err := NewGrid(g.rows, g.cols)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			neighbors := g.countNeighbors(i, j)
			alive := g.grid[i][j]

			if alive && (neighbors < 2 || neighbors > 3) {
				nextGen.SetCell(i, j, false)
			} else if !alive && neighbors == 3 {
				nextGen.SetCell(i, j, true)
			} else {
				nextGen.SetCell(i, j, alive)
			}
		}
	}

	return nextGen
}

// countNeighbors counts the neighbors of a cell.
func (g *Grid) countNeighbors(x, y int) int {
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
