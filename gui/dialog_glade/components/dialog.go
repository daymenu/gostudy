package components

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

// ShowDialog showdialog
func ShowDialog(content string, okFunc interface{}, cancelFunc interface{}) {
	log.Println(okFunc)
	dialogBuilder := getBuilder("ui/dialog.ui")
	dialog := getDialogFromBuilder(dialogBuilder, "dialog_glade_main")
	dialogLabel := getLabelFromBuilder(dialogBuilder, "dialog_glade_label")
	dialogLabel.SetLabel(content)
	dialog.ShowAll()
	buttonYes := getButtonFromBuilder(dialogBuilder, "dialog_glade_yes")
	buttonYes.Connect("clicked", func() {
		if call, ok := okFunc.(func()); ok {
			log.Println("components clicked")
			call()
		} else {
			log.Println("不是个函数瓦")
		}
		dialog.Destroy()
	})
	buttonNo := getButtonFromBuilder(dialogBuilder, "dialog_glade_no")
	buttonNo.Connect("clicked", func() {
		if call, ok := okFunc.(func()); ok {
			log.Println("components cancel clicked")
			call()
		}
		dialog.Destroy()
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
