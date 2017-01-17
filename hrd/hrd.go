package hrd

import (
	"fmt"
	"os"
	//"strconv"

	"github.com/nsf/termbox-go"
)

type Hrd struct {
	//Scorp    array
	W        int
	H        int
	Win      bool
	GameOver bool
	Exit     bool
	On       *Character
}

var charMap = map[string]*Character{
	"cc": &CC,
	"gy": &GY,
	"zf": &ZF,
	"zy": &ZY,
	"hz": &HZ,
	"mc": &MC,
	"s1": &S1,
	"s2": &S2,
	"s3": &S3,
	"s4": &S4,
	//"  ":&None,
}

var typeMap = map[string]string{
	"zf": "vertical",
	"zy": "vertical",
	"hz": "vertical",
	"mc": "vertical",
	"gy": "transverse",
	"cc": "square",
	"s1": "single",
	"s2": "single",
	"s3": "single",
	"s4": "single",
	"  ": "  ",
}

func (h *Hrd) Run() {
	h.Init()
	for {
		h.Loop()
		if h.Exit {
			return
		}
	}

}

func (h *Hrd) Init() {
	h.On = &S3
	h.On.On = true
	h.W = 0
	h.H = 4
	Draw()
}

func Draw() {
	DrawSurface()
	DrawElement()
	termbox.Flush()
}

func (h *Hrd) Loop() {
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			//			case termbox.KeyCtrlS:
			//				debug(a)
			//			case termbox.KeyCtrlA:
			//				debug1(strconv.Itoa(h.H) + " " + strconv.Itoa(h.W))
			case termbox.KeyCtrlR:
				break loop
			case termbox.KeyCtrlQ:
				h.Exit = true
				break loop
			case termbox.KeySpace:
				h.On.Choose = true
				PrintChar(h.On)
			case termbox.KeyArrowUp:
				if h.On.Choose == false {
					Up(h)
				} else {
					MoveUp(h)
					Up(h)
				}

			case termbox.KeyArrowDown:
				if h.On.Choose == false {
					Down(h)
				} else {
					MoveDown(h)
					Down(h)

				}

			case termbox.KeyArrowLeft:
				if h.On.Choose == false {
					Left(h)
				} else {
					MoveLeft(h)
					Left(h)
				}
			case termbox.KeyArrowRight:
				if h.On.Choose == false {
					Right(h)
				} else {
					MoveRight(h)
					Right(h)
				}
			}
		}
	}
}

func MoveDown(h *Hrd) {
	if h.H == 4 {
		return
	}
	switch typeMap[a[h.H][h.W]] {
	case "vertical":
		if h.H < 3 && a[h.H+2][h.W] == "  " {
			h.On.H = h.On.H + 4
			termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
			Draw()
			return
		}
	case "transverse":
		if h.H < 4 && a[h.H+1][h.W] == "  " && a[h.H+1][h.W+1] == "  " {
			h.On.H = h.On.H + 4
			termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
			Draw()
			return
		}
	case "square":
		if h.H < 3 && a[h.H+2][h.W] == "  " && a[h.H+2][h.W+1] == "  " {
			h.On.H = h.On.H + 4
			termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
			Draw()
			return
		}
	case "single":
		if h.H < 4 && a[h.H+1][h.W] == "  " {
			h.On.H = h.On.H + 4
			termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
			Draw()
			return
		}
	}
}

func MoveUp(h *Hrd) {
	if h.H == 0 {
		return
	}
	switch typeMap[a[h.H][h.W]] {
	case "vertical":
		if h.H > 0 && a[h.H-1][h.W] == "  " {
			h.On.H = h.On.H - 4
			termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
			Draw()
			return
		}
	case "transverse":
		if h.H > 0 && a[h.H-1][h.W] == "  " && a[h.H-1][h.W+1] == "  " {
			h.On.H = h.On.H - 4
			termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
			Draw()
			return
		}
	case "square":
		if h.H > 0 && a[h.H-1][h.W] == "  " && a[h.H-1][h.W+1] == "  " {
			h.On.H = h.On.H - 4
			termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
			Draw()
			return
		}
	case "single":
		if h.H > 0 && a[h.H-1][h.W] == "  " {
			h.On.H = h.On.H - 4
			termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
			Draw()
			return
		}
	}
}

func MoveLeft(h *Hrd) {
	switch typeMap[a[h.H][h.W]] {
	case "vertical":
		if h.W > 0 && a[h.H][h.W-1] == "  " && a[h.H+1][h.W-1] == "  " {
			h.On.W = h.On.W - 10
			termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
			Draw()
			return
		}
	case "transverse":
		if h.W > 0 && a[h.H][h.W-1] == "  " {
			h.On.W = h.On.W - 10
			termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
			Draw()
			return
		}
	case "square":
		if h.W > 0 && a[h.H][h.W-1] == "  " && a[h.H+1][h.W-1] == "  " {
			h.On.W = h.On.W - 10
			termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
			Draw()
			return
		}
	case "single":
		if h.W > 0 && a[h.H][h.W-1] == "  " {
			h.On.W = h.On.W - 10
			termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
			Draw()
			return
		}
	}
}

func MoveRight(h *Hrd) {
	switch typeMap[a[h.H][h.W]] {
	case "vertical":
		if h.W < 3 && a[h.H][h.W+1] == "  " && a[h.H+1][h.W+1] == "  " {
			h.On.W = h.On.W + 10
			termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
			Draw()
			return
		}
	case "transverse":
		if h.W < 2 && a[h.H][h.W+2] == "  " {
			h.On.W = h.On.W + 10
			termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
			Draw()
			return
		}
	case "square":
		if h.W < 2 && a[h.H][h.W+2] == "  " && a[h.H+1][h.W+2] == "  " {
			h.On.W = h.On.W + 10
			termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
			Draw()
			return
		}
	case "single":
		if h.W < 3 && a[h.H][h.W+1] == "  " {
			h.On.W = h.On.W + 10
			termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
			Draw()

			return
		}
	}
	h.On.Choose = false
	PrintChar(h.On)

}

func Up(h *Hrd) {
	if h.H == 0 {
		return
	}

	if a[h.H-1][h.W] == "  " {
		if h.On.Choose == true {
			ups(1, h)
			h.On.Choose = false
			PrintChar(h.On)
			return
		} else {
			return
		}
	}

	changeCharUpDown(-1, h)
	ups(1, h)

}

func ups(k int, h *Hrd) {
	b := changeArrayUp(k, h)

	switch typeMap[a[h.H-1][h.W]] {
	case "vertical":
		h.H = h.H - 2
	case "transverse":
		if h.W > 0 && a[h.H-1][h.W] == a[h.H-1][h.W-1] {
			h.H = h.H - 1
			h.W = h.W - 1
		} else {
			h.H = h.H - 1
		}
	case "square":
		if h.W > 0 && a[h.H-2][h.W] == a[h.H-2][h.W-1] {
			h.H = h.H - 2
			h.W = h.W - 1
		} else {
			h.H = h.H - 2
		}
	case "single":
		h.H = h.H - 1
	case "  ":
		h.H = h.H - 1
	}
	a = b
}

func Down(h *Hrd) {
	if h.H == 4 {
		return
	}
	switch typeMap[a[h.H][h.W]] {
	case "vertical":

		if h.H > 2 {
			return
		}
		if a[h.H+2][h.W] == "  " { /**/
			if h.On.Choose == true {
				//changeCharUpDown(2, h)
				downs(2, h)
				h.On.Choose = false
				PrintChar(h.On)
				return
			} else {
				return
			}
		}
		changeCharUpDown(2, h)
		downs(2, h)

	case "transverse":
		if a[h.H+1][h.W] == "  " { /**/
			if h.On.Choose == true {
				//changeCharUpDown(1, h)
				downs(1, h)
				h.On.Choose = false
				PrintChar(h.On)
				return
			} else {
				return
			}
		}
		changeCharUpDown(1, h)
		downs(1, h)
	case "square":
		if h.H > 2 {
			return
		}
		if a[h.H+2][h.W] == "  " { /**/
			if h.On.Choose == true {
				//changeCharUpDown(2, h)
				downs(2, h)
				h.On.Choose = false
				PrintChar(h.On)
				return
			} else {
				return
			}
		}
		changeCharUpDown(2, h)
		downs(2, h)
	case "single":
		if a[h.H+1][h.W] == "  " { /**/
			if h.On.Choose == true {
				//changeCharUpDown(1, h)
				downs(1, h)
				h.On.Choose = false
				PrintChar(h.On)
				return
			} else {
				return
			}
		}

		changeCharUpDown(1, h)
		downs(1, h)
	}

}

func Right(h *Hrd) {
	if h.W == 3 {
		return
	}

	switch typeMap[a[h.H][h.W]] {
	case "vertical":
		if a[h.H][h.W+1] == "  " { /**/
			if h.On.Choose == true {
				//changeCharLeftRight(1, h)
				rights(1, h)
				h.On.Choose = false
				PrintChar(h.On)
				return
			} else {
				return
			}
		}
		changeCharLeftRight(1, h)
		rights(1, h)
	case "transverse":
		if h.W > 1 {
			return
		}
		if a[h.H][h.W+2] == "  " { /**/
			if h.On.Choose == true {
				//changeCharLeftRight(2, h)
				rights(2, h)
				h.On.Choose = false
				PrintChar(h.On)
				return
			} else {
				return
			}
		}
		changeCharLeftRight(2, h)
		rights(2, h)
	case "square":
		if h.W > 1 {
			return
		}
		if a[h.H][h.W+2] == "  " { /**/
			if h.On.Choose == true {
				//changeCharLeftRight(2, h)
				rights(2, h)
				h.On.Choose = false
				PrintChar(h.On)
				return
			} else {
				return
			}
		}
		changeCharLeftRight(2, h)
		rights(2, h)
	case "single":
		if a[h.H][h.W+1] == "  " { /**/
			if h.On.Choose == true {
				//changeCharLeftRight(1, h)
				rights(1, h)
				h.On.Choose = false
				PrintChar(h.On)
				return
			} else {
				return
			}
		}
		changeCharLeftRight(1, h)
		rights(1, h)
	}
}

func Left(h *Hrd) {
	if h.W == 0 {
		return
	}
	if a[h.H][h.W-1] == "  " { /**/
		if h.On.Choose == true {
			//changeCharLeftRight(-1, h)
			lefts(h)
			h.On.Choose = false
			PrintChar(h.On)
			return
		} else {
			return
		}
	}
	changeCharLeftRight(-1, h)
	lefts(h)
}

func lefts(h *Hrd) {
	b := changeArrayLeft(-1, h)
	switch typeMap[a[h.H][h.W-1]] {
	case "vertical":
		if h.H > 0 && a[h.H][h.W-1] == a[h.H-1][h.W-1] {
			h.H = h.H - 1
			h.W = h.W - 1
		} else {
			h.W = h.W - 1
		}
	case "transverse":
		h.W = h.W - 2

	case "square":
		if h.H > 0 && a[h.H][h.W-1] == a[h.H-1][h.W-1] {
			h.H = h.H - 1
			h.W = h.W - 2
		} else {
			h.W = h.W - 2
		}
	case "single":
		h.W = h.W - 1
	case "  ":
		h.W = h.W - 1
	}
	a = b
}

func rights(k int, h *Hrd) {
	b := changeArrayRight(k, h)

	switch typeMap[a[h.H][h.W+k]] {
	case "vertical":
		if h.H > 0 && a[h.H][h.W+k] == a[h.H-1][h.W+k] {
			h.H = h.H - 1
			h.W = h.W + k
		} else {
			h.W = h.W + k
		}
	case "transverse":
		h.W = h.W + k

	case "square":
		if h.H > 0 && a[h.H][h.W+k] == a[h.H-1][h.W+k] {
			h.H = h.H - 1
			h.W = h.W + k
		} else {
			h.W = h.W + k
		}
	case "single":
		h.W = h.W + k
	case "  ":
		h.W = h.W + 1
	}
	a = b
}

func changeCharUpDown(k int, h *Hrd) {

	before := h.On
	before.On = false
	before.Choose = false
	h.On = charMap[a[h.H+k][h.W]]
	h.On.On = true
	PrintChar(before, h.On)
}

func changeCharLeftRight(k int, h *Hrd) {

	before := h.On
	before.On = false
	before.Choose = false
	h.On = charMap[a[h.H][h.W+k]]
	h.On.On = true
	PrintChar(before, h.On)
}

func downs(k int, h *Hrd) {
	b := changeArrayDown(k, h)
	switch typeMap[a[h.H+k][h.W]] {
	case "vertical":
		h.H = h.H + k
	case "transverse":
		if h.W > 0 && a[h.H+k][h.W] == a[h.H+k][h.W-1] {
			h.H = h.H + k
			h.W = h.W - 1
		} else {
			h.H = h.H + k
		}

	case "square":
		if h.W > 0 && a[h.H+k][h.W] == a[h.H+k][h.W-1] {
			h.H = h.H + k
			h.W = h.W - 1
		} else {
			h.H = h.H + k
		}
	case "single":
		h.H = h.H + k
	case "  ":
		h.H = h.H + 1
	}
	a = b
}

var a = [5][4]string{
	{"zf", "cc", "cc", "hz"},
	{"zf", "cc", "cc", "hz"},
	{"zy", "gy", "gy", "mc"},
	{"zy", "s1", "s2", "mc"},
	{"s3", "  ", "  ", "s4"},
}

func kkkkk() {
	fmt.Println(11111)
}

func changeArrayRight(k int, h *Hrd) (b [5][4]string) {
	b = a

	if h.On.Choose == true && b[h.H][h.W+k] == "  " {
		switch typeMap[b[h.H][h.W]] {

		case "vertical":
			if h.W < 3 && h.H < 4 && b[h.H][h.W+1] == "  " && b[h.H+1][h.W+1] == "  " {
				m, n := b[h.H][h.W], b[h.H+1][h.W]
				b[h.H][h.W] = b[h.H][h.W+1]
				b[h.H+1][h.W] = b[h.H+1][h.W+1]
				b[h.H][h.W+1] = m
				b[h.H+1][h.W+1] = n
			}
		case "transverse":
			if h.W < 2 && b[h.H][h.W+2] == "  " {
				m := b[h.H][h.W]
				b[h.H][h.W] = b[h.H][h.W+2]
				b[h.H][h.W+2] = m
			}
		case "square":
			if h.W < 2 && h.H < 4 && b[h.H][h.W+2] == "  " && b[h.H+1][h.W+2] == "  " {
				m := b[h.H][h.W]
				n := b[h.H+1][h.W]
				b[h.H][h.W] = b[h.H][h.W+2]
				b[h.H+1][h.W] = b[h.H+1][h.W+2]
				b[h.H][h.W+2] = m
				b[h.H+1][h.W+2] = n
			}
		case "single":
			m := b[h.H][h.W]
			b[h.H][h.W] = b[h.H][h.W+1]
			b[h.H][h.W+1] = m
		}
	}
	return
}

func changeArrayLeft(k int, h *Hrd) (b [5][4]string) {
	b = a
	if h.On.Choose == true && b[h.H][h.W+k] == "  " {
		switch typeMap[b[h.H][h.W]] {

		case "vertical":
			if h.W > 0 && h.H < 4 && b[h.H][h.W-1] == "  " && b[h.H+1][h.W-1] == "  " {
				m := b[h.H][h.W]
				b[h.H][h.W] = b[h.H][h.W-1]
				b[h.H+1][h.W] = b[h.H+1][h.W-1]
				b[h.H][h.W-1] = m
				b[h.H+1][h.W-1] = m
			}
		case "transverse":
			if h.W > 0 && b[h.H][h.W-1] == "  " {
				m := b[h.H][h.W]
				b[h.H][h.W+1] = b[h.H][h.W-1]
				b[h.H][h.W-1] = m
			}
		case "square":
			if h.W > 0 && h.H < 4 && b[h.H][h.W-1] == "  " && b[h.H+1][h.W-1] == "  " {
				m := b[h.H][h.W]
				b[h.H][h.W+1] = b[h.H][h.W-1]
				b[h.H+1][h.W+1] = b[h.H+1][h.W-1]
				b[h.H][h.W-1] = m
				b[h.H+1][h.W-1] = m
			}
		case "single":
			m := b[h.H][h.W]
			b[h.H][h.W] = b[h.H][h.W-1]
			b[h.H][h.W-1] = m
		}
	}
	return
}

func changeArrayUp(k int, h *Hrd) (b [5][4]string) {
	b = a
	if h.On.Choose == true && b[h.H-1][h.W] == "  " {
		switch typeMap[b[h.H][h.W]] {
		case "vertical":
			m := b[h.H][h.W]
			b[h.H+1][h.W] = b[h.H-1][h.W]
			b[h.H-1][h.W] = m
		case "transverse":
			if h.H > 0 && h.W < 3 && b[h.H-1][h.W] == "  " && b[h.H-1][h.W+1] == "  " {
				m := b[h.H][h.W]
				b[h.H][h.W] = b[h.H-1][h.W]
				b[h.H][h.W+1] = b[h.H-1][h.W+1]
				b[h.H-1][h.W] = m
				b[h.H-1][h.W+1] = m
			}
		case "square":
			if h.H > 0 && h.W < 3 && b[h.H-1][h.W] == "  " && b[h.H-1][h.W+1] == "  " {
				m := b[h.H][h.W]
				b[h.H+1][h.W] = b[h.H-1][h.W]
				b[h.H+1][h.W+1] = b[h.H-1][h.W+1]
				b[h.H-1][h.W] = m
				b[h.H-1][h.W+1] = m
			}

		case "single":
			m := b[h.H][h.W]
			b[h.H][h.W] = b[h.H-1][h.W]
			b[h.H-1][h.W] = m
		}
	}
	return
}

func changeArrayDown(k int, h *Hrd) (b [5][4]string) {
	b = a

	if h.On.Choose == true {

		switch typeMap[b[h.H][h.W]] {
		case "vertical":
			if h.H < 3 && b[h.H+2][h.W] == "  " {
				m := b[h.H][h.W]
				b[h.H][h.W] = b[h.H+2][h.W]
				b[h.H+2][h.W] = m
			}
		case "transverse":
			if h.H < 4 && h.W < 3 && b[h.H+1][h.W] == "  " && b[h.H+1][h.W+1] == "  " {
				m := b[h.H][h.W]
				b[h.H][h.W] = b[h.H+1][h.W]
				b[h.H][h.W+1] = b[h.H+1][h.W+1]
				b[h.H+1][h.W] = m
				b[h.H+1][h.W+1] = m
			}
		case "square":
			if h.H < 3 && h.W < 3 && b[h.H+2][h.W] == "  " && b[h.H+2][h.W+1] == "  " {
				m := b[h.H][h.W]
				b[h.H][h.W] = b[h.H+2][h.W]
				b[h.H][h.W+1] = b[h.H+2][h.W+1]
				b[h.H+2][h.W] = m
				b[h.H+2][h.W+1] = m
			}
		case "single":
			m := b[h.H][h.W]
			b[h.H][h.W] = b[h.H+1][h.W]
			b[h.H+1][h.W] = m
		}
	}
	return
}

func debug(s [5][4]string) {
	f, err := os.OpenFile("./b.text", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)

	if err != nil {
		fmt.Println(err)
		return
	}
	var str string
	for i := 0; i < 5; i++ {
		str = str + s[i][0] + s[i][1] + s[i][2] + s[i][3] + "\r\n"
	}

	_, err = f.WriteString(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Close()

}

func debug1(l string) {
	f, err := os.OpenFile("./a.text", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = f.WriteString(l + "\r\n")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Close()

}
