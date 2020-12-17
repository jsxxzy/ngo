package main

import (
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var inTE, outTE *walk.TextEdit

	var App = MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{320, 320},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					ListBox{},
					// TextEdit{AssignTo: &inTE},
					// TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "安装",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}

	App.Run()

}
