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
		"save": func (args []string) string {
			// format: save [filepath]
			if len(args) < 2 {
				return "Usage: save <filename>"
			}
			
			err := cv.Save(args[1])
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
			
			err := cv.Load(args[1])
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
			cv.DrawRect(x, y, w, h, char)
			
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
			cv.DrawLine(x1, y1, x2, y2, char)
			
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
			
			cv.DrawBox(x, y, w, h)
			
			return fmt.Sprintf("Box at (%d,%d) %dx%d", x, y, w, h)
		},
	}
}
