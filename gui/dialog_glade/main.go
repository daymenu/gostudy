package main

import (
	"log"
	"os"

	"github.com/daymenu/study/gui/dialog_glade/components"
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
		builder := getBuilder("ui/app.ui")
		setRapper(application, builder)
		windowObj, err := builder.GetObject("dialog_glade_app")
		if err != nil {
			log.Fatal("get main window faild:", err)
		}

		var window *gtk.ApplicationWindow
		var ok bool
		if window, ok = windowObj.(*gtk.ApplicationWindow); !ok {
			log.Fatal("dialog_glade_app is now window faild:", err)
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

func setRapper(app *gtk.Application, builder *gtk.Builder) {
	rapperButton := getButtonFromBuilder(builder, "dialog_glade_button_rapper")
	dialogBuilder := getBuilder("ui/dialog.ui")
	dialog := getDialogFromBuilder(dialogBuilder, "dialog_glade_main")
	rapperButton.Connect("clicked", func() {
		components.ShowDialog("你喜欢我吗？", func() {
			label := getLabelFromBuilder(builder, "dialog_glade_label")
			label.SetLabel("我喜欢你")
		}, func() {
			label := getLabelFromBuilder(builder, "dialog_glade_label")
			label.SetLabel("我不喜欢你")
		})
		log.Println("rapper is clicked")
		dialog.ShowAll()
		buttonYes := getButtonFromBuilder(dialogBuilder, "dialog_glade_yes")
		buttonYes.Connect("clicked", func() {
			label := getLabelFromBuilder(builder, "dialog_glade_label")
			label.SetLabel("您是个不折不扣的rapper")
			dialog.Destroy()
		})
		buttonNo := getButtonFromBuilder(dialogBuilder, "dialog_glade_no")
		buttonNo.Connect("clicked", func() {
			label := getLabelFromBuilder(builder, "dialog_glade_label")
			label.SetLabel("您不是个rapper")
			dialog.Destroy()
		})
	})
}

func getBuilder(uiPath string) *gtk.Builder {
	builder, err := gtk.BuilderNewFromFile(uiPath)
	if err != nil {
		log.Fatal("get main window faild:", err)
	}
	return builder
}

func getDialogFromBuilder(builder *gtk.Builder, id string) *gtk.Dialog {
	obj, err := builder.GetObject(id)
	if err != nil {
		log.Fatal("get "+id+" is not button:", err)
	}
	var dialog *gtk.Dialog
	var ok bool
	if dialog, ok = obj.(*gtk.Dialog); !ok {
		log.Fatal("get "+id+" is not dialog:", err)
	}
	return dialog
}

func getButtonFromBuilder(builder *gtk.Builder, id string) *gtk.Button {
	obj, err := builder.GetObject(id)
	if err != nil {
		log.Fatal("get dialog_glade_button_rapper is not button:", err)
	}
	var button *gtk.Button
	var ok bool
	if button, ok = obj.(*gtk.Button); !ok {
		log.Fatal("get dialog_glade_button_rapper is not button:", err)
	}
	return button
}

func getLabelFromBuilder(builder *gtk.Builder, id string) *gtk.Label {
	obj, err := builder.GetObject(id)
	if err != nil {
		log.Fatal("get dialog_glade_button_rapper is not button:", err)
	}
	var label *gtk.Label
	var ok bool
	if label, ok = obj.(*gtk.Label); !ok {
		log.Fatal("get dialog_glade_button_rapper is not button:", err)
	}
	return label
}
