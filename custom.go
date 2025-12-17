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

		"save": func(args []string) string {
			// format: save <filepath>
			if len(args) < 2 {
				return "Usage: save <filename>"
			}
			if err := cv.Save(args[1]); err != nil {
				return fmt.Sprintf("Error: %v", err)
			}
			return fmt.Sprintf("Saved to %s", args[1])
		},

		"load": func(args []string) string {
			// format: load <filepath>
			if len(args) < 2 {
				return "Usage: load <filename>"
			}
			if err := cv.Load(args[1]); err != nil {
				return fmt.Sprintf("Error: %v", err)
			}
			return fmt.Sprintf("Loaded from %s", args[1])
		},

		"fill": func(args []string) string {
			// format: fill <text>
			if len(args) < 2 {
				return "Usage: fill <text>"
			}
			cv.Fill(args[1])
			return fmt.Sprintf("Filled with '%s'", args[1])
		},

		"clear": func([]string) string {
			// format: clear
			for y := range cv.Data {
				for x := range cv.Data[y] {
					cv.Data[y][x] = ' '
				}
			}
			return "Canvas cleared"
		},

		"rect": func(args []string) string {
			// format: rect <x> <y> <w> <h> <char>
			if len(args) < 6 {
				return "Usage: rect <x> <y> <w> <h> <char>"
			}

			x, e1 := strconv.Atoi(args[1])
			y, e2 := strconv.Atoi(args[2])
			w, e3 := strconv.Atoi(args[3])
			h, e4 := strconv.Atoi(args[4])
			if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
				return "Error: invalid number format"
			}

			r := []rune(args[5])
			if len(r) != 1 {
				return "Error: char must be single character"
			}

			cv.DrawRect(x, y, w, h, r[0])
			return fmt.Sprintf("Rectangle at (%d,%d) %dx%d", x, y, w, h)
		},

		"line": func(args []string) string {
			// format: line <x1> <y1> <x2> <y2> <char>
			if len(args) < 6 {
				return "Usage: line <x1> <y1> <x2> <y2> <char>"
			}

			x1, e1 := strconv.Atoi(args[1])
			y1, e2 := strconv.Atoi(args[2])
			x2, e3 := strconv.Atoi(args[3])
			y2, e4 := strconv.Atoi(args[4])
			if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
				return "Error: invalid number format"
			}

			r := []rune(args[5])
			if len(r) != 1 {
				return "Error: char must be single character"
			}

			cv.DrawLine(x1, y1, x2, y2, r[0])
			return fmt.Sprintf("Line from (%d,%d) to (%d,%d)", x1, y1, x2, y2)
		},

		"box": func(args []string) string {
			// format: box <x> <y> <w> <h>
			if len(args) < 5 {
				return "Usage: box <x> <y> <w> <h>"
			}

			x, e1 := strconv.Atoi(args[1])
			y, e2 := strconv.Atoi(args[2])
			w, e3 := strconv.Atoi(args[3])
			h, e4 := strconv.Atoi(args[4])
			if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
				return "Error: invalid number format"
			}

			cv.DrawBox(x, y, w, h)
			return fmt.Sprintf("Box at (%d,%d) %dx%d", x, y, w, h)
		},

		"replace": func(args []string) string {
			// format: replace <old> <new> <count>
			if len(args) < 4 {
				return "Usage: replace <old> <new> <count>"
			}

			oldRune := []rune(args[1])
			newRune := []rune(args[2])
			if len(oldRune) != 1 || len(newRune) != 1 {
				return "Error: old and new must be single characters"
			}

			count, err := strconv.Atoi(args[3])
			if err != nil {
				return "Error: invalid replace count"
			}

			replaced := 0

			for y := 0; y < len(cv.Data); y++ {
				for x := 0; x < len(cv.Data[y]); x++ {
					if cv.Data[y][x] == oldRune[0] {
						cv.Data[y][x] = newRune[0]
						replaced++

						if count != -1 && replaced >= count {
							return fmt.Sprintf(
								"Replaced %d '%c' with '%c'",
								replaced, oldRune[0], newRune[0],
							)
						}
					}
				}
			}

			return fmt.Sprintf(
				"Replaced %d '%c' with '%c'",
				replaced, oldRune[0], newRune[0],
			)
		},

		"help": func([]string) string {
			return "" +
				"help - show this message\n" +
				"save <file> - save canvas to file\n" +
				"load <file> - load canvas from file\n" +
				"fill <text> - fill canvas with text\n" +
				"clear - clear the canvas\n" +
				"rect <x> <y> <w> <h> <c> - draw rectangle\n" +
				"line <x1> <y1> <x2> <y2> <c> - draw line\n" +
				"box <x> <y> <w> <h> - draw bordered box\n" +
				"replace <old> <new> <count> - replace characters (-1 = all)"
		},
	}
}
