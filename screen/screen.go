package screen

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sirupsen/logrus"
)

func Init() tcell.Screen {
	s, err := tcell.NewScreen()
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if err := s.Init(); err != nil {
		logrus.Error(err)
		return nil
	}
	s.HideCursor()
	return s
}

func Finish(s tcell.Screen) {
	s.Fini()
}

func Size(s tcell.Screen) (int, int) {
	return s.Size()
}
