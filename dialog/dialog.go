package dialog

type Dialog struct {
	Visible bool
	Prompt  string
	Buffer  []rune
}

func New() *Dialog {
	return &Dialog{}
}

func (d *Dialog) Show(prompt string) {
	d.Prompt = prompt
	d.Buffer = make([]rune, 0)
	d.Visible = true
}

func (d *Dialog) Hide() {
	d.Visible = false
	d.Prompt = ""
	d.Buffer = d.Buffer[:0]
}

func (d *Dialog) InputRune(r rune) {
	d.Buffer = append(d.Buffer, r)
}

func (d *Dialog) Backspace() {
	if len(d.Buffer) > 0 {
		d.Buffer = d.Buffer[:len(d.Buffer)-1]
	}
}

func (d *Dialog) GetInput() string {
	return string(d.Buffer)
}
