package main

import (
	"encoding/json"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type Config map[string]string

func main() {
	a := app.New()
	w := a.NewWindow("Dynon CSV Remapper")

	var inputPath, outputPath string
	config := loadConfig()

	inputBtn := widget.NewButton("Select Input Dynon CSV", func() {
		dialog.ShowFileOpen(func(f fyne.URIReadCloser, err error) {
			if err == nil && f != nil {
				inputPath = f.URI().Path()
				f.Close()
			}
		}, w)
	})

	outputBtn := widget.NewButton("Select Output Location", func() {
		dialog.ShowFileSave(func(f fyne.URIWriteCloser, err error) {
			if err == nil && f != nil {
				outputPath = f.URI().Path()
				f.Close()
			}
		}, w)
	})

	processBtn := widget.NewButton("Process", func() {
		if inputPath == "" || outputPath == "" {
			dialog.ShowInformation("Missing", "Select both input and output", w)
			return
		}
		// TODO: read CSV, apply config map, write output
		fmt.Println("Config loaded:", config)
		dialog.ShowInformation("Done", "Processing complete (stub)", w)
	})

	w.SetContent(container.NewVBox(
		widget.NewLabel("Dynon EMS → Savvy Aviation CSV Remapper"),
		inputBtn,
		outputBtn,
		processBtn,
	))
	w.ShowAndRun()
}

func loadConfig() Config {
	data, _ := os.ReadFile("config.json")
	var c Config
	json.Unmarshal(data, &c)
	return c
}
