package main

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Dice() {

	w := fyne.CurrentApp().NewWindow("Dice Game")
	r, _ := fyne.LoadResourceFromPath("D:\\Pep\\Image\\dice.png")
	w.SetIcon(r)
	w.CenterOnScreen()
	fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())

	w.Resize(fyne.NewSize(300, 200))

	img := canvas.NewImageFromFile("D:\\Pep\\Image\\12.png") // image of Dice
	img.FillMode = canvas.ImageFillOriginal

	btn1 := widget.NewButton("Play", func() {

		rand := rand.Intn(6) + 1 // ignore zero 0+1=1
		img.File = fmt.Sprintf("D:\\Pep\\Image\\%d.png", rand)
		img.Refresh()
	})

	w.SetContent(

		container.NewVBox(
			img,
			btn1,
		),
	)

	w.Show()
}
