package utils

import (
	"fmt"
)

/*
// ClearTerminal terminalの画面クリア

	func ClearTerminal() {
		switch runtime.GOOS {
		case "windows":
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
		default:
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	}
*/

// ClerTerminal terminalの画面クリア(ANSIエスケープコード版)
func ClearTerminal() {
	fmt.Print("\033[H\033[2J") // ANSIエスケープコード
}

// WaitForEnter ユーザの入力待ち
func WaitForEnter() {
	fmt.Print("Press Enter to continue...")
	fmt.Scanln()
}
