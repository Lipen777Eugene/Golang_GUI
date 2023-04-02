package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("title")
	w.Resize(fyne.NewSize(400, 400))
	//w.SetContent(widget.NewLabel("here is first Label"))

	check2 := widget.NewCheck("Check", func(b bool) {
		if b {
			fmt.Println("This true")
		} else {
			fmt.Println("This false")
		}
	})
	w.SetContent(check2)
	w.ShowAndRun()
}
