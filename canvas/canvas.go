package canvas

type Canvas struct {
	Width  int
	Height int
	Data   [][]rune
	CX     int
	CY     int
}

func NewCanvas(w, h int) *Canvas {
	d := make([][]rune, h)
	for i := range d {
		d[i] = make([]rune, w)
		for j := range d[i] {
			d[i][j] = ' '
		}
	}
	return &Canvas{Width: w, Height: h, Data: d}
}

func (c *Canvas) PutChar(r rune) {
	if c.CX >= 0 && c.CX < c.Width && c.CY >= 0 && c.CY < c.Height {
		c.Data[c.CY][c.CX] = r
	}
	if c.CX+1 < c.Width {
		c.CX++
	}
}

func (c *Canvas) MoveCursor(dx, dy int) {
	nx := c.CX + dx
	ny := c.CY + dy
	if nx >= 0 && nx < c.Width {
		c.CX = nx
	}
	if ny >= 0 && ny < c.Height {
		c.CY = ny
	}
}

func (c *Canvas) Backspace() {
	if c.CX > 0 {
		c.CX--
		c.Data[c.CY][c.CX] = ' '
	}
}

func (c *Canvas) Enter() {
	if c.CY+1 < c.Height {
		c.CY++
		c.CX = 0
	}
}
