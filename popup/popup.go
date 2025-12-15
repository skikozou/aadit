package popup

type Popup struct {
	Visible bool
	Message string
}

func New() *Popup {
	return &Popup{}
}

func (p *Popup) Show(msg string) {
	p.Message = msg
	p.Visible = true
}

func (p *Popup) Hide() {
	p.Visible = false
	p.Message = ""
}
