package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type GuiObj struct {
	Win fyne.Window
}

func makeBanner() fyne.CanvasObject {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {}),
	)
	logo := canvas.NewImageFromResource(resourceLogoJpg)
	logo.FillMode = canvas.ImageFillContain
	return container.NewStack(toolbar, logo)
}

func makeLeftMenu() fyne.CanvasObject {
	return widget.NewLabel("Serial configuration")
}

func makeRightMenu() fyne.CanvasObject {
	return widget.NewLabel("Advanced Settings")
}

func (g *GuiObj) MakeGUI() fyne.CanvasObject {
	top := makeBanner()
	left := makeLeftMenu()
	right := makeRightMenu()

	content := canvas.NewRectangle(color.Gray{Y: 0xee})

	dividers := [3]fyne.CanvasObject{
		widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(),
	}
	objs := []fyne.CanvasObject{content, top, left, right, dividers[0], dividers[1], dividers[2]}
	return container.New(newHermesLayout(top, left, right, content, dividers), objs...)
}
