package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/skip2/go-qrcode"
)

func QrCode() {

	w := fyne.CurrentApp().NewWindow("Qr Code Generator")
	r, _ := fyne.LoadResourceFromPath("D:\\Pep\\Image\\qrcode.png")
	w.SetIcon(r)
	w.CenterOnScreen()

	fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())

	w.Resize(fyne.NewSize(400, 400))

	photo := canvas.NewImageFromFile("D:\\Pep\\Image\\qr.png")
	photo.FillMode = canvas.ImageFillOriginal

	// qrcode generator
	url := widget.NewEntry()
	url.SetPlaceHolder("Enter url ...")
	size := widget.NewEntry()
	size.SetPlaceHolder("Enter file size i.e 256 , etc.")
	size_1, _ := strconv.Atoi(size.Text)
	file_name := widget.NewEntry()
	file_name.SetPlaceHolder("Enter file name ...")

	btn := widget.NewButton("Create", func() {

		// quality / resolution
		// size of image 256x256
		c := qrcode.WriteFile(
			url.Text,
			qrcode.Highest,
			size_1,
			fmt.Sprintf("%s.png", file_name.Text), // sprintf will add suffix .png
		)
		if c != nil {
			fmt.Println(c)
		}
	})
	btn.Importance = widget.HighImportance

	w.SetContent(container.NewVBox(
		photo,
		url,
		size,
		file_name,
		btn,
	))
	w.Show()
}
