package main

import (
	"lifegame-go/game"
	"log"
	"math/rand"
	"time"
)

// Gridサイズ
const (
	width  = 50
	height = 30
)

func main() {
	rand.Seed(time.Now().UnixNano())

	grid, err := initializeGrid(height, width)
	if err != nil {
		log.Fatal(err)
		return
	}

	runGame(grid)
}

func initializeGrid(rows, cols int) (*game.Grid, error) {
	grid, err := game.NewGrid(rows, cols)
	if err != nil {
		return nil, err
	}

	// ランダムに初期化
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			grid.SetCell(i, j, rand.Intn(2) == 1)
		}
	}

	return grid, nil
}

func runGame(initialGrid *game.Grid) {
	grid := initialGrid

	for {
		grid.Display()
		time.Sleep(time.Second / 2)

		// 新しい世代の grid を取得
		grid = grid.NextGeneration()
	}
}
