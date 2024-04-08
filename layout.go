package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const sidewidth = 220

type fysionLayout struct {
	top, left, right, content fyne.CanvasObject
	dividers                  [3]fyne.CanvasObject
}

func newFysionLayout(top, left, right, content fyne.CanvasObject, dividers [3]fyne.CanvasObject) fyne.Layout {
	return &fysionLayout{top: top, left: left, right: right, content: content, dividers: dividers}
}

func (l *fysionLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	topHeight := l.top.MinSize().Height
	l.top.Resize(fyne.NewSize(size.Width, topHeight))

	l.left.Move(fyne.NewPos(0, topHeight))
	l.left.Resize(fyne.NewSize(sidewidth, size.Height-topHeight))

	l.right.Move(fyne.NewPos(size.Width-sidewidth, topHeight))
	l.right.Resize(fyne.NewSize(sidewidth, size.Height-topHeight))

	l.content.Move(fyne.NewPos(sidewidth, topHeight))
	l.content.Resize(fyne.NewSize(size.Width-sidewidth*2, size.Height-topHeight))

	dividerThickness := theme.SeparatorThicknessSize()
	l.dividers[0].Move(fyne.NewPos(0, topHeight))
	l.dividers[0].Resize(fyne.NewSize(size.Width, dividerThickness))

	l.dividers[1].Move(fyne.NewPos(sidewidth, topHeight))
	l.dividers[1].Resize(fyne.NewSize(dividerThickness, size.Height-topHeight))

	l.dividers[2].Move(fyne.NewPos(size.Width-sidewidth, topHeight))
	l.dividers[2].Resize(fyne.NewSize(dividerThickness, size.Height-topHeight))
}

func (l *fysionLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	borders := fyne.NewSize(
		sidewidth*2,
		l.top.MinSize().Height,
	)
	return borders.AddWidthHeight(200, 200)
}
