package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("title")

	w.Resize(fyne.NewSize(400, 400))

	colorx := color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	textX := canvas.NewText("here is my text", colorx)

	w.SetContent(textX)
	w.ShowAndRun()
}
