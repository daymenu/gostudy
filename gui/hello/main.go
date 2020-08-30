package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// 不解析命令行参数的初始化GTK
	gtk.Init(nil)

	// 创建最高层的Window窗口
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	// 设置窗口标题
	win.SetTitle("Simple Example")
	// 窗口结束，退出GTK
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// 创建一个文字为Hello, gotk3!的标签
	l, err := gtk.LabelNew("Hello, gotk3!")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	// 将标签加入窗口
	win.Add(l)

	// 设置窗口的默认大小
	win.SetDefaultSize(800, 600)

	// 显示窗口中的所有内容.
	win.ShowAll()

	// 开始运行GTK主循环，直到gtk.MainQuit()执行
	gtk.Main()
}
