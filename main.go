package main

import (
	"Hermes/gui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	mainApp := app.New()
	mainApp.Settings().SetTheme(gui.NewHermesTheme())
	mainWindow := mainApp.NewWindow("Hermes")
	mainWindow.Resize(fyne.NewSize(1024, 768))

	mainGui := gui.GuiObj{Win: mainWindow}
	mainWindow.SetContent(mainGui.MakeGUI())
	mainWindow.ShowAndRun()
}
