package main

import (
	"fmt"
	"os"
	"strconv"
	"aadit/canvas"
	"aadit/command"
	"aadit/dialog"
	"aadit/popup"
	
	"github.com/gdamore/tcell/v2"
)

func Customize(s tcell.Screen, cv *canvas.Canvas, con *command.Console, pop *popup.Popup, dlg *dialog.Dialog) {
	con.Functions = map[string]command.Function{
		"save": func (args []string) string {
			// format: save [filepath]
			if len(args) < 2 {
				return "Usage: save <filename>"
			}
			
			err := saveCanvas(cv, args[1])
			if err != nil {
				return fmt.Sprintf("Error: %v", err)
			}
			return fmt.Sprintf("Saved to %s", args[1])
		},

		"load": func (args []string) string {
			// format: load [filepath]
			if len(args) < 2 {
				return "Usage: load <filename>"
			}
			
			err := loadCanvas(cv, args[1])
			if err != nil {
				return fmt.Sprintf("Error: %v", err)
			}
			return fmt.Sprintf("Loaded from %s", args[1])
		},
		
 		"fill": func (args []string) string {
 			// format: fill [text]
		   	if len(args) < 2 {
				return "Usage: fill <text>"
		   	}
		    cv.Fill(args[1])
		    return fmt.Sprintf("Filled with '%s'", args[1])
	    },
		
		"help": func ([]string) string {
 			// format: help
			return "Available commands:\n" +
				"help - show this message\n" +
				"save <file> - save canvas to file\n" +
				"load <file> - load canvas from file\n" +
				"fill <text> - fill canvas with text\n" +
				"clear - clear the canvas\n" +
				"rect <x> <y> <w> <h> <c> - draw rectangle\n" +
				"line <x1> <y1> <x2> <y2> <c> - draw line\n" +
				"box <x> <y> <w> <h> - draw bordered box"
		},
		
		"clear": func ([]string) string {
 			// format: clear
			for i := range cv.Data {
				for j := range cv.Data[i] {
					cv.Data[i][j] = ' '
				}
			}
			return "Canvas cleared"
		},
		
		"rect": func (args []string) string {
 			// format: rect <x> <y> <w> <h> <char>
			if len(args) < 6 {
				return "Usage: rect <x> <y> <w> <h> <char>"
			}
			
			x, err1 := strconv.Atoi(args[1])
			y, err2 := strconv.Atoi(args[2])
			w, err3 := strconv.Atoi(args[3])
			h, err4 := strconv.Atoi(args[4])
			
			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				return "Error: invalid number format"
			}
			
			char := rune(args[5][0])
			drawRect(cv, x, y, w, h, char)
			
			return fmt.Sprintf("Rectangle at (%d,%d) %dx%d", x, y, w, h)
		},
		
		"line": func (args []string) string {
 			// format: line <x1> <y1> <x2> <y2> <char>
			if len(args) < 6 {
				return "Usage: line <x1> <y1> <x2> <y2> <char>"
			}
			
			x1, err1 := strconv.Atoi(args[1])
			y1, err2 := strconv.Atoi(args[2])
			x2, err3 := strconv.Atoi(args[3])
			y2, err4 := strconv.Atoi(args[4])
			
			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				return "Error: invalid number format"
			}
			
			char := rune(args[5][0])
			drawLine(cv, x1, y1, x2, y2, char)
			
			return fmt.Sprintf("Line from (%d,%d) to (%d,%d)", x1, y1, x2, y2)
		},
		
		"box": func (args []string) string {
 			// format: box <x> <y> <w> <h>
			if len(args) < 5 {
				return "Usage: box <x> <y> <w> <h>"
			}
			
			x, err1 := strconv.Atoi(args[1])
			y, err2 := strconv.Atoi(args[2])
			w, err3 := strconv.Atoi(args[3])
			h, err4 := strconv.Atoi(args[4])
			
			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				return "Error: invalid number format"
			}
			
			drawBox(cv, x, y, w, h)
			
			return fmt.Sprintf("Box at (%d,%d) %dx%d", x, y, w, h)
		},
	}
}

func saveCanvas(cv *canvas.Canvas, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	
	for _, row := range cv.Data {
		for _, ch := range row {
			if _, err := file.WriteString(string(ch)); err != nil {
				return err
			}
		}
		if _, err := file.WriteString("\n"); err != nil {
			return err
		}
	}
	
	return nil
}

func loadCanvas(cv *canvas.Canvas, filepath string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	
	lines := [][]rune{}
	currentLine := []rune{}
	
	for _, ch := range string(data) {
		if ch == '\n' {
			lines = append(lines, currentLine)
			currentLine = []rune{}
		} else {
			currentLine = append(currentLine, ch)
		}
	}
	
	// 最後の行が改行で終わっていない場合
	if len(currentLine) > 0 {
		lines = append(lines, currentLine)
	}
	
	// キャンバスに適用
	for i := 0; i < cv.Height && i < len(lines); i++ {
		for j := 0; j < cv.Width && j < len(lines[i]); j++ {
			cv.Data[i][j] = lines[i][j]
		}
		// 行が短い場合は空白で埋める
		for j := len(lines[i]); j < cv.Width; j++ {
			cv.Data[i][j] = ' '
		}
	}
	
	// ファイルの行数がキャンバスより少ない場合は空白で埋める
	for i := len(lines); i < cv.Height; i++ {
		for j := 0; j < cv.Width; j++ {
			cv.Data[i][j] = ' '
		}
	}
	
	return nil
}

func drawRect(cv *canvas.Canvas, x, y, w, h int, char rune) {
	for dy := 0; dy < h; dy++ {
		for dx := 0; dx < w; dx++ {
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < cv.Width && ny >= 0 && ny < cv.Height {
				cv.Data[ny][nx] = char
			}
		}
	}
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
