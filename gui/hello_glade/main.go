package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const appID = "github.com.dayment.gostudy.gui.hello_glade"

func main() {
	// 创建 application
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatal("create application faild:", err)
	}

	// applicatioin 激活的时候加载视图文件
	application.Connect("activate", func() {
		log.Println("application activate")
		builder, err := gtk.BuilderNewFromFile("glade/hello.glade")

		if err != nil {
			log.Fatal("create builder faild:", err)
		}

		windowObj, err := builder.GetObject("hello_glade_main")
		if err != nil {
			log.Fatal("get main window faild:", err)
		}

		var window *gtk.ApplicationWindow
		var ok bool
		if window, ok = windowObj.(*gtk.ApplicationWindow); !ok {
			log.Fatal("hello_glade_main is now window faild:", err)
		}

		window.ShowAll()

		application.AddWindow(window)
	})

	// application 结束时执行的操作
	application.Connect("shutdown", func() {
		log.Println("application shutdown")
	})

	// 运行 application
	os.Exit(application.Run(os.Args))
}
