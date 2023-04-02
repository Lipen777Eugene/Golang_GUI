package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("title")
	w.Resize(fyne.NewSize(400, 400))
	//w.SetContent(widget.NewLabel("here is first Label"))

	label1 := widget.NewLabel("I can write anything")
	w.SetContent(label1)

	w.ShowAndRun()
}
