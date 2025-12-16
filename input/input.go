package input

import (
	"aadit/canvas"
	"aadit/command"
	"aadit/dialog"
	"aadit/popup"

	"github.com/gdamore/tcell/v2"
)

func HandleEvent(ev *tcell.EventKey, cv *canvas.Canvas, con *command.Console, pop *popup.Popup, dlg *dialog.Dialog) bool {

	// Ctrl + /
	if ev.Key() == tcell.KeyCtrlUnderscore {
		con.Toggle()
		return true
	}

	// ダイアログ表示中の処理
	if dlg.Visible {
		switch ev.Key() {
		case tcell.KeyEscape:
			dlg.Cancel()
		case tcell.KeyEnter:
			dlg.Submit()
		case tcell.KeyBackspace, tcell.KeyBackspace2:
			dlg.Backspace()
		case tcell.KeyRune:
			dlg.InputRune(ev.Rune())
		}
		return true
	}

	if pop.Visible {
		if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyEnter {
			pop.Hide()
		}
		return true
	}

	// コンソール優先
	if con.Visible {
		switch ev.Key() {
		case tcell.KeyEscape:
			con.Toggle()
		case tcell.KeyEnter:
			result := con.Execute()
			if result != "" {
				pop.Show(result)
			}
		case tcell.KeyBackspace, tcell.KeyBackspace2:
			con.Backspace()
		case tcell.KeyRune:
			con.InputRune(ev.Rune())
		}
		return true
	}

	// キャンバス操作
	switch ev.Key() {
	case tcell.KeyEscape:
		return false
	case tcell.KeyUp:
		cv.MoveCursor(0, -1)
	case tcell.KeyDown:
		cv.MoveCursor(0, 1)
	case tcell.KeyLeft:
		cv.MoveCursor(-1, 0)
	case tcell.KeyRight:
		cv.MoveCursor(1, 0)
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		cv.Backspace()
	case tcell.KeyEnter:
		cv.Enter()
	case tcell.KeyRune:
		cv.PutChar(ev.Rune())
	}
	return true
}
