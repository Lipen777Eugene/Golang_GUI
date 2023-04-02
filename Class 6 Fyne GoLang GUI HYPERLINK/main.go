package main

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Title")
	w.Resize(fyne.NewSize(400, 400))

	url1, _ := url.Parse("https://google.com/")

	fmt.Println(url1)
	hyperlink := widget.NewHyperlink("Visit me", url1)

	w.SetContent(hyperlink)
	w.ShowAndRun()
}
