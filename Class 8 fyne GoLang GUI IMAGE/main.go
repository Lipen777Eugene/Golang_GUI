package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("title")
	w.Resize(fyne.NewSize(450, 450))

	img := canvas.NewImageFromFile("./image.jpg")

	w.SetContent(img)
	w.ShowAndRun()
}
