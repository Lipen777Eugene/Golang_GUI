package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Title")
	w.Resize(fyne.NewSize(400, 400))

	btn := widget.NewButton("button Name", func() {
		fmt.Println("Here is Go Button")
	})
	w.SetContent(btn)
	w.ShowAndRun()
}
