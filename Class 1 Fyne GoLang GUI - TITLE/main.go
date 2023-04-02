package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	ap := app.New()
	window := ap.NewWindow("Hello World")

	window.ShowAndRun()
}
