package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func Audio() {

	var format beep.Format
	var streamer beep.StreamSeekCloser
	var pause bool = false

	//a := app.New()
	w := fyne.CurrentApp().NewWindow("Spotify")
	w.CenterOnScreen() //set position

	r, _ := fyne.LoadResourceFromPath("D:\\Pep\\Image\\spot.png")
	w.SetIcon(r)

	fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
	w.Resize(fyne.NewSize(500, 300))

	logo := canvas.NewImageFromFile("D:\\Pep\\Image\\spo2.png")
	logo.FillMode = canvas.ImageFillOriginal
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(), // this is to give space from left
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {

			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			speaker.Play(streamer)
		}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			if !pause {
				pause = true
				speaker.Lock()
			} else if pause {
				pause = false
				speaker.Unlock()
			}
		}),
		widget.NewToolbarAction(theme.MediaStopIcon(), func() {
			speaker.Clear()

		}),
		widget.NewToolbarSpacer(), // space from right
	)
	label1 := widget.NewLabel("Audio MP3..")
	label1.Alignment = fyne.TextAlignCenter
	label2 := widget.NewLabel("Play MP3..")
	label2.Alignment = fyne.TextAlignCenter

	browse_files := widget.NewButton("Browse...", func() {
		fd := dialog.NewFileOpen(func(uc fyne.URIReadCloser, _ error) { // newfileOPEN is pre defined syntax to open the drive/dialog box & uriREADcloser to read that file
			streamer, format, _ = mp3.Decode(uc)
			label2.Text = uc.URI().Name()
			label2.Refresh()
		}, w)
		fd.Show()
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))
	})

	c := container.NewVBox(logo, label1, browse_files, label2, toolbar)
	w.SetContent(c)
	w.Show()
}
