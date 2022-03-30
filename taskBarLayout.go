package main

import (
	"fyne.io/fyne/v2"

)

type diagonal struct{}

func (d *diagonal) MinSize(items []fyne.CanvasObject) fyne.Size {
	total := fyne.NewSize(0, 0)
	for _,obj := range items {
		if !obj.Visible(){
			continue
		}
		total = total.Add(obj.MinSize())
	}
	return total
}

func (d *diagonal) Layout(items []fyne.CanvasObject, size fyne.Size)  {
	topLeft := fyne.NewPos(size.Width/2-200, size.Height-58)
	for _, obj := range items {
		if !obj.Visible(){
			continue
	}
	Mysize := fyne.NewSize(55, 55)

	obj.Move(topLeft)
	obj.Resize(Mysize)
	topLeft.X = topLeft.X + 70
}

}