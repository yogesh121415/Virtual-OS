package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

)

func main() {
	app := app.New()
	win := app.NewWindow("Mac OS ")

	icon, _ := fyne.LoadResourceFromPath("D:\\Pep\\Image\\apple.jpg")
	win.SetIcon(icon)

	win.CenterOnScreen()
	win.SetPadded(false)
	win.Resize(fyne.NewSize(1000, 600))
	wallpaper := canvas.NewImageFromFile("D:\\Pep\\Image\\bg10.png")


	win.SetContent(container.NewBorder(nil, nil, nil, nil, wallpaper,TaskbarIconUI()))
	win.ShowAndRun()
}
