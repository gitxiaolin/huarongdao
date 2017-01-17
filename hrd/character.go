package hrd

import (
	"github.com/nsf/termbox-go"
)

type Character struct {
	W      int
	H      int
	Weight int
	Height int
	Name1  rune
	Name2  rune
	Color  termbox.Attribute
	On     bool
	Choose bool
}

var CC, GY, ZF, ZY, HZ, MC, S1, S2, S3, S4, None Character

func init() {
	CC = Character{20, 0, 20, 8, 'c', 'c', termbox.ColorBlue, false, false}
	GY = Character{20, 8, 20, 4, 'g', 'y', termbox.ColorRed, false, false}
	ZF = Character{10, 0, 10, 8, 'z', 'f', termbox.ColorCyan, false, false}
	ZY = Character{10, 8, 10, 8, 'z', 'y', termbox.ColorCyan, false, false}
	HZ = Character{40, 0, 10, 8, 'h', 'z', termbox.ColorCyan, false, false}
	MC = Character{40, 8, 10, 8, 'm', 'c', termbox.ColorCyan, false, false}
	S1 = Character{20, 12, 10, 4, 's', '1', termbox.ColorCyan, false, false}
	S2 = Character{30, 12, 10, 4, 's', '2', termbox.ColorCyan, false, false}
	S3 = Character{10, 16, 10, 4, 's', '3', termbox.ColorCyan, false, false}
	S4 = Character{40, 16, 10, 4, 's', '4', termbox.ColorCyan, false, false}
}
