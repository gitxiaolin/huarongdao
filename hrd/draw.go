package hrd

import (
	"github.com/nsf/termbox-go"
)

func DrawSurface() {
	// 绘制边框
	termbox.SetCell(10, 0, 0x250c, termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(50, 0, 0x2510, termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(10, 20, 0x2514, termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(50, 20, 0x2518, termbox.ColorYellow, termbox.ColorDefault)

	for i := 11; i < 50; i++ {
		termbox.SetCell(i, 0, 0x2500, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(i, 20, 0x2500, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(i, 4, 0x2500, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(i, 8, 0x2500, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(i, 12, 0x2500, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(i, 16, 0x2500, termbox.ColorYellow, termbox.ColorDefault)
	}

	for i := 1; i < 20; i++ {
		termbox.SetCell(10, i, 0x2502, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(20, i, 0x2502, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(30, i, 0x2502, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(40, i, 0x2502, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(50, i, 0x2502, termbox.ColorYellow, termbox.ColorDefault)
	}

	termbox.Flush()
}

func DrawElement() {
	PrintChar(&CC, &GY, &ZF, &ZY, &HZ, &MC, &S1, &S2, &S3, &S4)
}

func PrintChar(v ...*Character) {
	for _, c := range v {

		for i := c.H + 1; i < c.H+c.Height; i++ {
			for j := c.W + 1; j < c.W+c.Weight; j++ {
				if j == c.W+2 && i == c.H+2 {
					termbox.SetCell(j, i, c.Name1, termbox.ColorYellow, termbox.ColorDefault)
					continue
				}
				if j == c.W+3 && i == c.H+2 {
					termbox.SetCell(j, i, c.Name2, termbox.ColorYellow, termbox.ColorDefault)
					continue
				}
				termbox.SetCell(j, i, ' ', termbox.ColorYellow, c.Color)
			}
		}
		if c.Choose {
			termbox.SetCell(c.W+2, c.H+2, c.Name1, termbox.ColorYellow, termbox.ColorBlack)
			termbox.SetCell(c.W+3, c.H+2, c.Name2, termbox.ColorYellow, termbox.ColorBlack)
		} else if c.On {
			termbox.SetCell(c.W+2, c.H+2, c.Name1, termbox.ColorWhite, termbox.ColorBlack)
			termbox.SetCell(c.W+3, c.H+2, c.Name2, termbox.ColorWhite, termbox.ColorBlack)
		}
	}
	termbox.Flush()
}
