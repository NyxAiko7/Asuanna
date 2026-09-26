package main

import (
	//"fmt"
	//"os"
	//"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
  	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func main() {

	user := loadConfig()

	a := app.New()
	w := a.NewWindow("programa swap wallpeper")
	w.Resize(fyne.NewSize(1920, 1080))

	// Метка
	label := widget.NewLabel("programa swap wallpeper")

	//wallpaper wvwie
	img := canvas.NewImageFromFile(user.Path)
	img.FillMode = canvas.ImageFillContain
	//img.SetMinSize(fyne.NewSize(400, 300))

	// Кнопка с обработчиком нажатия
	button := widget.NewButton("restart ags", func() {
		restartags(user)
	})

	button1 := widget.NewButton("pywal", func(){
		getpywal(user)
	})

	button2 := widget.NewButton("telegram", func(){
		getpywalTelegram(user)
	})

	button3 := widget.NewButton("install wallpeper",func(){
		installwallppepar(user)
	})

	button4 := widget.NewButton("restart swaync", func(){
		restarswaync(user)
	})

	top := container.NewVBox(
		label,
		button,
	 	button1,
	 	button2,
	 	button3,
	 	button4,
	)
	w.SetContent(container.NewBorder(top, nil, nil, nil, img))


	w.ShowAndRun()
}
