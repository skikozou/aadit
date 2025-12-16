// ===================
//       Waring!
// ===================
//
// This system is not supported
//

package dialog

type Dialog struct {
	Visible  bool
	Prompt   string
	Buffer   []rune
	Callback func(string)
}

func New() *Dialog {
	return &Dialog{}
}

func (d *Dialog) Show(prompt string, callback func(string)) {
	d.Prompt = prompt
	d.Buffer = make([]rune, 0)
	d.Callback = callback
	d.Visible = true
}

func (d *Dialog) Hide() {
	d.Visible = false
	d.Prompt = ""
	d.Buffer = d.Buffer[:0]
	d.Callback = nil
}

func (d *Dialog) InputRune(r rune) {
	d.Buffer = append(d.Buffer, r)
}

func (d *Dialog) Backspace() {
	if len(d.Buffer) > 0 {
		d.Buffer = d.Buffer[:len(d.Buffer)-1]
	}
}

func (d *Dialog) Submit() {
	if d.Callback != nil {
		d.Callback(string(d.Buffer))
	}
	d.Hide()
}

func (d *Dialog) Cancel() {
	d.Hide()
}
