package draw

import (
	"fmt"

	"aadit/canvas"
	"aadit/command"
	"aadit/popup"
	"aadit/screen"

	"github.com/gdamore/tcell/v2"
)

func DrawAll(s tcell.Screen, cv *canvas.Canvas, con *command.Console, pop *popup.Popup) {
	s.Clear()

	sw, sh := screen.Size(s)
	cx := (sw - cv.Width) / 2
	cy := (sh - cv.Height) / 2

	// ---- 枠 ----
	for x := 0; x < cv.Width+2; x++ {
		s.SetContent(cx-1+x, cy-1, '─', nil, tcell.StyleDefault)
		s.SetContent(cx-1+x, cy+cv.Height, '─', nil, tcell.StyleDefault)
	}
	for y := 0; y < cv.Height; y++ {
		s.SetContent(cx-1, cy+y, '│', nil, tcell.StyleDefault)
		s.SetContent(cx+cv.Width, cy+y, '│', nil, tcell.StyleDefault)
	}
	s.SetContent(cx-1, cy-1, '┌', nil, tcell.StyleDefault)
	s.SetContent(cx+cv.Width, cy-1, '┐', nil, tcell.StyleDefault)
	s.SetContent(cx-1, cy+cv.Height, '└', nil, tcell.StyleDefault)
	s.SetContent(cx+cv.Width, cy+cv.Height, '┘', nil, tcell.StyleDefault)

	// ---- 本体 ----
	for y := 0; y < cv.Height; y++ {
		for x := 0; x < cv.Width; x++ {
			style := tcell.StyleDefault
			if x == cv.CX && y == cv.CY {
				style = style.Reverse(true)
			}
			s.SetContent(cx+x, cy+y, cv.Data[y][x], nil, style)
		}
	}

	// ---- コンソール（白背景・黒文字） ----
	if con.Visible {
		style := tcell.StyleDefault.
			Background(tcell.ColorWhite).
			Foreground(tcell.ColorBlack)

		y := cy + cv.Height + 1
		s.SetContent(cx-1, y, '>', nil, style)

		for i := 0; i < cv.Width; i++ {
			ch := ' '
			if i < len(con.Buffer) {
				ch = con.Buffer[i]
			}
			s.SetContent(cx+i, y, ch, nil, style)
		}
	}

	// ---- ポップアップ ----
	if pop.Visible {
		w := len(pop.Message) + 4
		h := 3
		px := (sw - w) / 2
		py := (sh - h) / 2

		style := tcell.StyleDefault.
			Background(tcell.ColorWhite).
			Foreground(tcell.ColorBlack)

		// 枠
		for x := 0; x < w; x++ {
			s.SetContent(px+x, py, '─', nil, style)
			s.SetContent(px+x, py+h-1, '─', nil, style)
		}
		for y := 0; y < h; y++ {
			s.SetContent(px, py+y, '│', nil, style)
			s.SetContent(px+w-1, py+y, '│', nil, style)
		}
		s.SetContent(px, py, '┌', nil, style)
		s.SetContent(px+w-1, py, '┐', nil, style)
		s.SetContent(px, py+h-1, '└', nil, style)
		s.SetContent(px+w-1, py+h-1, '┘', nil, style)

		s.SetContent(px+1, py+1, ' ', nil, style)
		s.SetContent(px+w-2, py+1, ' ', nil, style)
		for i, r := range pop.Message {
			s.SetContent(px+2+i, py+1, r, nil, style)
		}
	}

	// 左下
	for i, r := range "AAdit" {
		s.SetContent(i, sh-1, r, nil, tcell.StyleDefault)
	}

	// 右下（キャンバスサイズ）
	rs := fmt.Sprintf("%dx%d", cv.Width, cv.Height)
	for i, r := range rs {
		s.SetContent(sw-len(rs)+i, sh-1, r, nil, tcell.StyleDefault)
	}

	s.Show()
}
