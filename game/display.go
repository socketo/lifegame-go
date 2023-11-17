package game

import (
	"fmt"
	"lifegame-go/utils"
)

// Display グリッドの表示
func (g *Grid) Display() {
	utils.ClearTerminal()

	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			if g.grid[i][j] {
				fmt.Print("\x1b[40m  \x1b[0m") // 黒
			} else {
				fmt.Print("\x1b[47m  \x1b[0m") // 白
			}
		}
		fmt.Print("\n")
	}

	fmt.Println()
}
