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

	rect := canvas.NewRectangle(color.Black)

	rect.FillColor = color.NRGBA{R: 255, G: 255, B: 0, A: 255}
	rect.StrokeColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	rect.StrokeWidth = 5.0
	w.SetContent(rect)
	w.ShowAndRun()
}
