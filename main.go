package main

import (
        "os"
        "strconv"
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

        cw, _ := strconv.Atoi(os.Args[1])
        ch, _ := strconv.Atoi(os.Args[2])

        s := screen.Init()
        if s == nil {
                return
        }
        defer screen.Finish(s)

        cv := canvas.NewCanvas(cw, ch)
        con := command.NewConsole()
        pop := popup.New()

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
