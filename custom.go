package main

import (
	"fmt"
	"aadit/command"
	"aadit/canvas"
	"aadit/popup"
	"aadit/dialog"
	
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
			return "help - show this message\nfill [text] - fill canvas by text"
		},
		"dlg": func ([]string) string {
			
		
			return ""
		},
	}
}
