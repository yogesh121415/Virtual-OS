package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	//"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Weather2() {


	w:= fyne.CurrentApp().NewWindow("Weather")
	r, _ := fyne.LoadResourceFromPath("D:\\Pep\\Image\\weather4.jpg")
	w.SetIcon(r)
	fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
	w.CenterOnScreen(

	)
	w.Resize(fyne.NewSize(500,500))
	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=noida&APPID=5959337e2386cc70c9c084a1d6f28179")
	if err != nil {
		fmt.Print(err)
	}
	defer res.Body.Close()

	info, _ := ioutil.ReadAll(res.Body)

	weather, _ := UnmarshalWeather(info) // Used api.QuickType.io to Convert Json response

	// Now we are almost done
	img := canvas.NewImageFromFile("D:\\Pep\\Image\\weather4.jpg")

	img.SetMinSize(fyne.NewSize(300, 500))

	label2 := canvas.NewText(fmt.Sprintf("Country - %s", weather.Sys.Country), color.White)
	label2.TextStyle = fyne.TextStyle{Bold: true}
	label5 := canvas.NewText(fmt.Sprintf("City - %s", weather.Name), color.White)
	label5.TextStyle = fyne.TextStyle{Bold: true}
	label3 := canvas.NewText(fmt.Sprintf("Wind Speed - %.2f mps", weather.Wind.Speed), color.White)
	label3.TextStyle = fyne.TextStyle{Bold: true}
	label4 := canvas.NewText(fmt.Sprintf("Temperature - %.2f K", weather.Main.Temp), color.White)
	label4.TextStyle = fyne.TextStyle{Bold: true}
	label7 := canvas.NewText(fmt.Sprintf("Humidity - %d RH", weather.Main.Humidity), color.White)
	label7.TextStyle = fyne.TextStyle{Bold: true}

	break3 := widget.NewLabel(" ")

	w.SetContent(
		container.NewBorder(nil, nil, nil, nil, img, container.NewCenter(container.NewVBox(
			label2,
			label5,
			break3,
			label3,
			label4,
			label7,
		))),
	)
	w.Resize(fyne.NewSize(400, 400))
	w.SetPadded(false)
	w.Show() // show and run  the app
}

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}
func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`
	Weather    []WeatherElement `json:"weather"`
	Base       string           `json:"base"`
	Main       Main             `json:"main"`
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`
	Clouds     Clouds           `json:"clouds"`
	Dt         int64            `json:"dt"`
	Sys        Sys              `json:"sys"`
	Timezone   int64            `json:"timezone"`
	ID         int64            `json:"id"`
	Name       string           `json:"name"`
	Cod        int64            `json:"cod"`
}
type Clouds struct {
	All int64 `json:"all"`
}
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}
type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
	SeaLevel  int64   `json:"sea_level"`
	GrndLevel int64   `json:"grnd_level"`
}
type Sys struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}
type WeatherElement struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
	Gust  float64 `json:"gust"`
}
