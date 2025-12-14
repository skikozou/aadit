package draw

import (
	"fmt"

	"aadit/canvas"
	"aadit/screen"

	"github.com/gdamore/tcell/v2"
)

func DrawAll(s tcell.Screen, cv *canvas.Canvas) {
	s.Clear()

	sw, sh := screen.Size(s)

	cx := (sw - cv.Width) / 2
	cy := (sh - cv.Height) / 2

	// ----------- キャンバス枠 --------------
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

	// ----------- 本体 --------------
	for y := 0; y < cv.Height; y++ {
		for x := 0; x < cv.Width; x++ {
			ch := cv.Data[y][x]
			style := tcell.StyleDefault

			// カーソル反転
			if x == cv.CX && y == cv.CY {
				style = style.Reverse(true)
			}

			s.SetContent(cx+x, cy+y, ch, nil, style)
		}
	}

	// ----------- 右下：キャンバスサイズ --------------
	rs := fmt.Sprintf("%dx%d", cv.Width, cv.Height)
	for i, r := range rs {
		s.SetContent(sw-len(rs)+i, sh-1, r, nil, tcell.StyleDefault)
	}

	// ----------- 左下：AAdit --------------
	a := "AAdit"
	for i, r := range a {
		s.SetContent(i, sh-1, r, nil, tcell.StyleDefault)
	}

	s.Show()
}
