package main

import (
	"fmt"
	"strconv"
	"aadit/canvas"
	"aadit/command"
	"aadit/dialog"
	"aadit/popup"
	
	"github.com/gdamore/tcell/v2"
)

func Customize(s tcell.Screen, cv *canvas.Canvas, con *command.Console, pop *popup.Popup, dlg *dialog.Dialog) {
	con.Functions = map[string]command.Function{
 		"fill": func (args []string) string {
		   	if len(args) < 2 {
		   		return "missing args"
		   	}
		    cv.Fill(args[1])
		    return fmt.Sprintf("Filled %s", args[1])
	    },
		
		"help": func ([]string) string {
			return "help - show this message\nfill [text] - fill canvas with repeating text\nclear - clear the canvas\nrect - draw a rectangle\nline - draw a line\nbox - draw a box with borders"
		},
		
		"clear": func ([]string) string {
			for i := range cv.Data {
				for j := range cv.Data[i] {
					cv.Data[i][j] = ' '
				}
			}
			return "Canvas cleared"
		},
		
		"rect": func ([]string) string {
			dlg.Show("x y width height char:", func(input string) {
				parts := splitInput(input)
				if len(parts) < 5 {
					pop.Show("Error: need 5 arguments (x y width height char)")
					return
				}
				
				x, err1 := strconv.Atoi(parts[0])
				y, err2 := strconv.Atoi(parts[1])
				w, err3 := strconv.Atoi(parts[2])
				h, err4 := strconv.Atoi(parts[3])
				
				if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
					pop.Show("Error: invalid number format")
					return
				}
				
				char := rune(parts[4][0])
				
				for dy := 0; dy < h; dy++ {
					for dx := 0; dx < w; dx++ {
						nx, ny := x+dx, y+dy
						if nx >= 0 && nx < cv.Width && ny >= 0 && ny < cv.Height {
							cv.Data[ny][nx] = char
						}
					}
				}
				
				pop.Show(fmt.Sprintf("Drew rectangle at (%d,%d) size %dx%d", x, y, w, h))
			})
			return ""
		},
		
		"line": func ([]string) string {
			dlg.Show("x1 y1 x2 y2 char:", func(input string) {
				parts := splitInput(input)
				if len(parts) < 5 {
					pop.Show("Error: need 5 arguments (x1 y1 x2 y2 char)")
					return
				}
				
				x1, err1 := strconv.Atoi(parts[0])
				y1, err2 := strconv.Atoi(parts[1])
				x2, err3 := strconv.Atoi(parts[2])
				y2, err4 := strconv.Atoi(parts[3])
				
				if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
					pop.Show("Error: invalid number format")
					return
				}
				
				char := rune(parts[4][0])
				
				drawLine(cv, x1, y1, x2, y2, char)
				
				pop.Show(fmt.Sprintf("Drew line from (%d,%d) to (%d,%d)", x1, y1, x2, y2))
			})
			return ""
		},
		
		"box": func ([]string) string {
			dlg.Show("x y width height:", func(input string) {
				parts := splitInput(input)
				if len(parts) < 4 {
					pop.Show("Error: need 4 arguments (x y width height)")
					return
				}
				
				x, err1 := strconv.Atoi(parts[0])
				y, err2 := strconv.Atoi(parts[1])
				w, err3 := strconv.Atoi(parts[2])
				h, err4 := strconv.Atoi(parts[3])
				
				if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
					pop.Show("Error: invalid number format")
					return
				}
				
				drawBox(cv, x, y, w, h)
				
				pop.Show(fmt.Sprintf("Drew box at (%d,%d) size %dx%d", x, y, w, h))
			})
			return ""
		},
	}
}

func splitInput(input string) []string {
	var parts []string
	var current string
	
	for _, ch := range input {
		if ch == ' ' {
			if current != "" {
				parts = append(parts, current)
				current = ""
			}
		} else {
			current += string(ch)
		}
	}
	
	if current != "" {
		parts = append(parts, current)
	}
	
	return parts
}

func drawLine(cv *canvas.Canvas, x1, y1, x2, y2 int, char rune) {
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	sx := sign(x2 - x1)
	sy := sign(y2 - y1)
	err := dx - dy
	
	x, y := x1, y1
	
	for {
		if x >= 0 && x < cv.Width && y >= 0 && y < cv.Height {
			cv.Data[y][x] = char
		}
		
		if x == x2 && y == y2 {
			break
		}
		
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x += sx
		}
		if e2 < dx {
			err += dx
			y += sy
		}
	}
}

func drawBox(cv *canvas.Canvas, x, y, w, h int) {
	if w < 2 || h < 2 {
		return
	}
	
	// 上下
	for dx := 0; dx < w; dx++ {
		nx := x + dx
		if nx >= 0 && nx < cv.Width {
			if y >= 0 && y < cv.Height {
				if dx == 0 {
					cv.Data[y][nx] = '┌'
				} else if dx == w-1 {
					cv.Data[y][nx] = '┐'
				} else {
					cv.Data[y][nx] = '─'
				}
			}
			if y+h-1 >= 0 && y+h-1 < cv.Height {
				if dx == 0 {
					cv.Data[y+h-1][nx] = '└'
				} else if dx == w-1 {
					cv.Data[y+h-1][nx] = '┘'
				} else {
					cv.Data[y+h-1][nx] = '─'
				}
			}
		}
	}
	
	// 左右
	for dy := 1; dy < h-1; dy++ {
		ny := y + dy
		if ny >= 0 && ny < cv.Height {
			if x >= 0 && x < cv.Width {
				cv.Data[ny][x] = '│'
			}
			if x+w-1 >= 0 && x+w-1 < cv.Width {
				cv.Data[ny][x+w-1] = '│'
			}
		}
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sign(n int) int {
	if n < 0 {
		return -1
	}
	if n > 0 {
		return 1
	}
	return 0
}
