package main

import (
        "os"
        "fmt"
        "strconv"
        "strings"
        "aadit/canvas"
        "aadit/draw"
        "aadit/input"
        "aadit/logger"
        "aadit/screen"
        "aadit/command"
        "aadit/popup"

        "github.com/gdamore/tcell/v2"
)

func main() {
        logger.LogrusInit()

		var cw, ch int
        switch len(os.Args)-1 {
        	case 0:
        		// format: none
        		cw, ch = 0, 0
        	case 1:
        		// format: 00x00
        		size := strings.Split(os.Args[1], "x")
        		if len(size) < 2 {
        			cw, ch = 0, 0
        		} else {
        			width, err1 := strconv.Atoi(size[0])
        			height, err2 := strconv.Atoi(size[1])
        			if err1 != nil || err2 != nil {
        				cw, ch = 0, 0
        			} else {
        				cw, ch = width, height
        			}
        		}
        	case 2:
        		// format: 00 00
        		width, err1 := strconv.Atoi(os.Args[1])
        		height, err2 := strconv.Atoi(os.Args[2])
       			if err1 != nil || err2 != nil {
       				cw, ch = 0, 0
       			} else {
       				cw, ch = width, height
       			}
       		default:
       			// format: other
       			cw, ch = 0, 0
        }

        s := screen.Init()
        if s == nil {
                return
        }
        defer screen.Finish(s)

        if cw*ch == 0 {
        	width, height := s.Size()
        	cw, ch = width/2, height/2
        }

        cv := canvas.NewCanvas(cw, ch)
        con := command.NewConsole()
        pop := popup.New()

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
        }

        for {
                draw.DrawAll(s, cv, con, pop)

                ev := s.PollEvent()
                switch e := ev.(type) {
                case *tcell.EventKey:
                        if !input.HandleEvent(e, cv, con, pop) {
                                return
                        }
                }
        }
}
