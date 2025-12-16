package canvas

import "os"

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
	c.Data[c.CY][c.CX] = r
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
	c.Data[c.CY][c.CX] = ' '
}

func (c *Canvas) Enter() {
	if c.CY+1 < c.Height {
		c.CY++
		c.CX = 0
	}
}

func (c *Canvas) Fill(text string) {
	for i := range c.Data {
		for j := range c.Data[i] {
			c.Data[i][j] = rune(text[(i*c.Width+j)%(len(text))])
		}
	}
}

func (c *Canvas)DrawRect(x, y, w, h int, char rune) {
	for dy := 0; dy < h; dy++ {
		for dx := 0; dx < w; dx++ {
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < c.Width && ny >= 0 && ny < c.Height {
				c.Data[ny][nx] = char
			}
		}
	}
}

func (c *Canvas)DrawLine(x1, y1, x2, y2 int, char rune) {
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	sx := sign(x2 - x1)
	sy := sign(y2 - y1)
	err := dx - dy
	
	x, y := x1, y1
	
	for {
		if x >= 0 && x < c.Width && y >= 0 && y < c.Height {
			c.Data[y][x] = char
		}
		
		if x == x2 && y == y2 {
			break
		}
		
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x += sx
		}
		if e2 < dx {
			err += dx
			y += sy
		}
	}
}

func (c *Canvas)DrawBox(x, y, w, h int) {
	if w < 2 || h < 2 {
		return
	}
	
	// 上下
	for dx := 0; dx < w; dx++ {
		nx := x + dx
		if nx >= 0 && nx < c.Width {
			if y >= 0 && y < c.Height {
				if dx == 0 {
					c.Data[y][nx] = '┌'
				} else if dx == w-1 {
					c.Data[y][nx] = '┐'
				} else {
					c.Data[y][nx] = '─'
				}
			}
			if y+h-1 >= 0 && y+h-1 < c.Height {
				if dx == 0 {
					c.Data[y+h-1][nx] = '└'
				} else if dx == w-1 {
					c.Data[y+h-1][nx] = '┘'
				} else {
					c.Data[y+h-1][nx] = '─'
				}
			}
		}
	}
	
	// 左右
	for dy := 1; dy < h-1; dy++ {
		ny := y + dy
		if ny >= 0 && ny < c.Height {
			if x >= 0 && x < c.Width {
				c.Data[ny][x] = '│'
			}
			if x+w-1 >= 0 && x+w-1 < c.Width {
				c.Data[ny][x+w-1] = '│'
			}
		}
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sign(n int) int {
	if n < 0 {
		return -1
	}
	if n > 0 {
		return 1
	}
	return 0
}

func (c *Canvas)Save(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	
	for _, row := range c.Data {
		for _, ch := range row {
			if _, err := file.WriteString(string(ch)); err != nil {
				return err
			}
		}
		if _, err := file.WriteString("\n"); err != nil {
			return err
		}
	}
	
	return nil
}

func (c *Canvas)Load(filepath string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	
	lines := [][]rune{}
	currentLine := []rune{}
	
	for _, ch := range string(data) {
		if ch == '\n' {
			lines = append(lines, currentLine)
			currentLine = []rune{}
		} else {
			currentLine = append(currentLine, ch)
		}
	}
	
	// 最後の行が改行で終わっていない場合
	if len(currentLine) > 0 {
		lines = append(lines, currentLine)
	}
	
	// キャンバスに適用
	for i := 0; i < c.Height && i < len(lines); i++ {
		for j := 0; j < c.Width && j < len(lines[i]); j++ {
			c.Data[i][j] = lines[i][j]
		}
		// 行が短い場合は空白で埋める
		for j := len(lines[i]); j < c.Width; j++ {
			c.Data[i][j] = ' '
		}
	}
	
	// ファイルの行数がキャンバスより少ない場合は空白で埋める
	for i := len(lines); i < c.Height; i++ {
		for j := 0; j < c.Width; j++ {
			c.Data[i][j] = ' '
		}
	}
	
	return nil
}
