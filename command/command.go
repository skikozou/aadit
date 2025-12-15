package command

import (
	"fmt"
	"github.com/google/shlex"
	)

type Console struct {
	Visible bool
	Buffer  []rune
	Functions map[string]Function
}

type Function func([]string) string

func NewConsole() *Console {
	return &Console{
		Buffer: make([]rune, 0),
		Functions: map[string]Function{
					"fill": func (args []string) string {
					        	if len(args) < 2 {
					        		return "missing args"
					        	}
					
					        	return fmt.Sprintf("Filled %s", args[1])
					        },
					"help": func ([]string) string {
						return "help message\nmaybe"
					},
				},
	}
}

func (c *Console) Toggle() {
	c.Visible = !c.Visible
	c.Buffer = c.Buffer[:0]
}

func (c *Console) InputRune(r rune) {
	c.Buffer = append(c.Buffer, r)
}

func (c *Console) Backspace() {
	if len(c.Buffer) > 0 {
		c.Buffer = c.Buffer[:len(c.Buffer)-1]
	}
}

func (c *Console) Execute() string {
	parts, _ := shlex.Split(string(c.Buffer))
	c.Buffer = c.Buffer[:0]
	var message string

	if len(parts) < 1 {
		return "empty command"
	}

	cmdFunc := c.Functions[parts[0]]
	if cmdFunc == nil {
		return "unkown command: " + parts[0]
	}

	message = cmdFunc(parts)

	return message
}
