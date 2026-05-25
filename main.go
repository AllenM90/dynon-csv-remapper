package main

import (
	"fmt"
	"os"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type Config map[string]string

func main() {
	var mw *walk.MainWindow
	var inputPath, outputPath string
	config := loadConfig()

	MainWindow{
		AssignTo: &mw,
		Title:    "Dynon CSV Remapper",
		Size:     Size{300, 200},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text: "Dynon EMS → Savvy Aviation CSV Remapper",
			},
			PushButton{
				Text: "Select Input Dynon CSV",
				OnClicked: func() {
					dlg := new(walk.FileDialog)
					dlg.Title = "Select Input CSV"
					if ok, _ := dlg.ShowOpen(mw); ok {
						inputPath = dlg.FilePath
					}
				},
			},
			PushButton{
				Text: "Select Output Location",
				OnClicked: func() {
					dlg := new(walk.FileDialog)
					dlg.Title = "Select Output CSV"
					if ok, _ := dlg.ShowSave(mw); ok {
						outputPath = dlg.FilePath
					}
				},
			},
			PushButton{
				Text: "Process",
				OnClicked: func() {
					if inputPath == "" || outputPath == "" {
						walk.MsgBox(mw, "Missing", "Select both input and output", walk.MsgBoxIconWarning)
						return
					}
					fmt.Println("Config loaded:", config)
					walk.MsgBox(mw, "Done", "Processing complete (stub)", walk.MsgBoxIconInformation)
				},
			},
		},
	}.Create()

	mw.Run()
}

func loadConfig() Config {
	data, _ := os.ReadFile("config.json")
	var c Config
	json.Unmarshal(data, &c)
	return c
}
