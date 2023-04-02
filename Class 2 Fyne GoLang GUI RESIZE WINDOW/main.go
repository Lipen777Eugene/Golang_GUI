package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	ap := app.New()
	window := ap.NewWindow("Title")
	window.Resize(fyne.NewSize(400, 400))
	window.ShowAndRun()
}
