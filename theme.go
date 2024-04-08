//go:generate fyne bundle -o bundled.go assets

package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type fysionTheme struct {
	fyne.Theme
}

func newFysionTheme() fyne.Theme {
	return &fysionTheme{Theme: theme.DefaultTheme()}
}

func (t *fysionTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return t.Theme.Color(name, theme.VariantLight)
}
