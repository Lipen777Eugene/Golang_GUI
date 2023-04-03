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

	circle1 := canvas.NewCircle(color.NRGBA{R: 0, G: 0, B: 255, A: 255})
	circle1.StrokeColor = color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	circle1.StrokeWidth = 10
	w.SetContent(circle1)
	w.ShowAndRun()
}
