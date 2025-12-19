package input

import (
	"aadit/canvas"
	"aadit/command"
	"aadit/dialog"
	"aadit/popup"

	"github.com/gdamore/tcell/v2"
)

func HandleMouseEvent(ev *tcell.EventMouse, cv *canvas.Canvas, con *command.Console, pop *popup.Popup, dlg *dialog.Dialog) bool {
	x, y := ev.Position()
	//btn := ev.Buttons()

	//on canvas
	if onCanvas(cv, x, y) {
		//nope
	}

	return false
}

func onCanvas(cv *canvas.Canvas, x, y int) bool {
	return false
}
