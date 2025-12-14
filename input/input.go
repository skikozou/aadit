package input

import (
	"aadit/canvas"

	"github.com/gdamore/tcell/v2"
)

func HandleEvent(ev *tcell.EventKey, cv *canvas.Canvas) bool {
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
