package main

import (
	"Hermes/gui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(gui.NewHermesTheme())
	w := a.NewWindow("Hermes")
	w.Resize(fyne.NewSize(1024, 768))
	w.SetContent(gui.MakeGUI())
	w.ShowAndRun()
}
