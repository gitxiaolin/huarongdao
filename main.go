package main

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"

	"huarongdao/hrd"
)

var h hrd.Hrd

func main() {
	fmt.Println("按上、下、左、右来移动光标，按空格选中模块，然后按上、下、左、右来移动模块。")
	fmt.Println("输入任意内容按回车开始")

	var buf string
	fmt.Scan(&buf)
	err := termbox.Init()
	if err != nil {
		fmt.Printf("初始化失败, Error:%s\n", err)
		os.Exit(1)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	h.Run()

}
