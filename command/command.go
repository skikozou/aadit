package command

type Console struct {
	Visible bool
	Buffer  []rune
}

func NewConsole() *Console {
	return &Console{
		Buffer: make([]rune, 0),
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
	cmd := string(c.Buffer)
	c.Buffer = c.Buffer[:0]

	if cmd == "" {
		return "empty command"
	}

	// 今はダミー
	return "executed: " + cmd
}
