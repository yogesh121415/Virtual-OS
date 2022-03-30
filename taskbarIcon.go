package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type tapIcon struct{
	widget.Icon // we dont need to give datatypes as fyne understands itself
 	tap    func()
}

// we are making custom tapable icon in taskbarIcon.go
func newTapIcon(res fyne.Resource, fn func()) *tapIcon {
	i := &tapIcon{tap: fn}
	i.Resource = res 
	i.ExtendBaseWidget(i)
	return i
}

func (t *tapIcon) Tapped(_ *fyne.PointEvent){
	t.tap()
}

func TaskbarIconUI() fyne.CanvasObject {
	//audio work

	res, _ := fyne.LoadResourceFromPath("D:\\Pep\\Image\\spot.png")
	audioImage := theme.NewThemedResource(res)
	app1 := newTapIcon(audioImage, func() {
		Audio()
	})

	//cal work
	res2, _ := fyne.LoadResourceFromPath("D:\\Pep\\Image\\calculator.png")
	calcImage := theme.NewThemedResource(res2)
	app2 := newTapIcon(calcImage, func(){
		Calc()
	})
	
	//photo

	res3, _ := fyne.LoadResourceFromPath("D:\\Pep\\Image\\gallery.png")
	galleryImage := theme.NewThemedResource(res3)
	app3 := newTapIcon(galleryImage, func(){
		Gallery()
	})
	//weather

	res4, _ := fyne.LoadResourceFromPath("D:\\Pep\\Image\\weather2.png")
	weatherImage := theme.NewThemedResource(res4)
	app4 := newTapIcon(weatherImage, func(){
		Weather2()
	})
//dice
//res5, _ := fyne.LoadResourceFromPath("D:\\Pep\\Image\\dice.png")
//diceImage := theme.NewThemedResource(res5)
//app5 := newTapIcon(diceImage, func(){
//	Dice()
//})

res6, _ := fyne.LoadResourceFromPath("D:\\Pep\\Image\\qrcode2.jpg")
qrcodeImage := theme.NewThemedResource(res6)
app6 := newTapIcon(qrcodeImage, func(){
	QrCode()
})

res7, _ := fyne.LoadResourceFromPath("D:\\Pep\\Image\\note.png")
textEditorImage := theme.NewThemedResource(res7)
app7 := newTapIcon(textEditorImage, func(){
	Notepad()
})


tasksbarContainer := container.New(&diagonal{},app1,app2,app3,app4,app6,app7)

return tasksbarContainer
}