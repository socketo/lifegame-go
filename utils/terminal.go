package utils

import (
	"fmt"
)

// ClerTerminal terminalの画面クリア(ANSIエスケープコード版)
func ClearTerminal() {
	fmt.Print("\033[H\033[2J") // ANSIエスケープコード
}

// WaitForEnter ユーザの入力待ち
func WaitForEnter() {
	fmt.Print("Press Enter to continue...")
	fmt.Scanln()
}
