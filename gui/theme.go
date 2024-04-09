//go:generate fyne bundle -o bundled.go assets

package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type HermesTheme struct {
	fyne.Theme
}

func NewHermesTheme() fyne.Theme {
	return &HermesTheme{Theme: theme.DefaultTheme()}
}

func (t *HermesTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return t.Theme.Color(name, theme.VariantLight)
}
