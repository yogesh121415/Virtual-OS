package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var count int = 1

func Notepad() {

	w := fyne.CurrentApp().NewWindow("Text Editor")
	r, _ := fyne.LoadResourceFromPath("D:\\Pep\\Image\\Notepad.png")
	w.SetIcon(r)
	w.CenterOnScreen()
	fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())

	w.Resize(fyne.NewSize(500, 500))

	content := container.NewVBox(

		widget.NewLabel("Pep Text Editor"),
	)
	content.Add(widget.NewButton("Add New File", func() {
		content.Add(widget.NewLabel("New File" + strconv.Itoa(count)))
		count++
	}))

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")
	input.Resize(fyne.NewSize(400, 400))

	saveBtn := widget.NewButton("Save text File", func() {
		saveFileDialog := dialog.NewFileSave(func(uc fyne.URIWriteCloser, _ error) { //this is callback fncn
			textData := []byte(input.Text) //content entry

			uc.Write(textData) //save the text in the textData
		}, w)
		saveFileDialog.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")

		saveFileDialog.Show()
	})

	openBtn := widget.NewButton("Open Text File", func() {
		openFIleDialog := dialog.NewFileOpen(func(r fyne.URIReadCloser, _ error) {

			ReadData, _ := ioutil.ReadAll(r) // we read all the data

			//go to NewWindow
			output := fyne.NewStaticResource("New File", ReadData)

			viewData := widget.NewMultiLineEntry()

			viewData.SetText(string(output.StaticContent))

			w := fyne.CurrentApp().NewWindow(
				string(output.StaticName))

			w.SetContent(container.NewScroll(viewData))
			w.Resize(fyne.NewSize(400, 400))
			w.Show()
		}, w)

		openFIleDialog.SetFilter(
			storage.NewExtensionFileFilter([]string{".txt"}))

		openFIleDialog.Show()
	})

	w.SetContent(
		container.NewVBox(
			content,
			input,

			container.NewHBox(
				saveBtn,
				openBtn,
			),
		),
	)
	w.Show()
}
