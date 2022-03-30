package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func Gallery() {

	w := fyne.CurrentApp().NewWindow("Gallery")
	r,_ := fyne.LoadResourceFromPath("D:\\Pep\\Image\\apple.png")
	w.SetIcon(r)
	w.CenterOnScreen()
	fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
	w.Resize(fyne.NewSize(500, 500))

	location := "D:\\Pep\\Gallery"

	files, _ := ioutil.ReadDir(location)

	tabs := container.NewAppTabs()

	for _, file := range files { // first dash is for index second is for range
		fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir() {
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "PNG" || extension == "jpg" {
				image := canvas.NewImageFromFile(location + "\\" + file.Name())
				image.FillMode = canvas.ImageFillOriginal
				tabs.Append(container.NewTabItem(file.Name(), image)) // apped all files satisfyiung this ciondition
			}

		}
	}

	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs)
	w.Show()
}
