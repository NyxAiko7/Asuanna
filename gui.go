package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	currentUser := loadConfig("")

	a := app.New()
	w := a.NewWindow("Asuanna — Wallpaper Switcher")
	w.Resize(fyne.NewSize(1280, 800))

	img := canvas.NewImageFromFile(currentUser.Path)
	img.FillMode = canvas.ImageFillContain

	title := widget.NewLabelWithStyle("Wallpaper Control", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	configStatus := widget.NewLabel("Config: Default")
	configStatus.TextStyle = fyne.TextStyle{Italic: true}

	btnSelectConfig := widget.NewButtonWithIcon("Select config.json", theme.FolderOpenIcon(), func() {
		fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil || reader == nil {
				return
			}
			defer reader.Close()

			selectedPath := reader.URI().Path()
			currentUser = loadConfig(selectedPath)

			configStatus.SetText("Config: " + reader.URI().Name())
			img.File = currentUser.Path
			img.Refresh()
		}, w)

		fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))
		fileDialog.Show()
	})

	btnInstall := widget.NewButtonWithIcon("Apply Wallpaper", theme.ConfirmIcon(), func() {
		installwallppepar(currentUser)
	})
	btnInstall.Importance = widget.HighImportance

	btnPywal := widget.NewButtonWithIcon("Pywal", theme.ColorPaletteIcon(), func() {
		getpywal(currentUser)
		dialog.ShowInformation("Pywal", "Pywal theme updated!", w)
	})
	btnTelegram := widget.NewButtonWithIcon("Telegram", theme.MailSendIcon(), func() {
		getpywalTelegram(currentUser)
		dialog.ShowInformation("Telegram", "Telegram theme applied!", w)
	})
	pywalGrid := container.NewGridWithColumns(2, btnPywal, btnTelegram)

	btnAgs := widget.NewButtonWithIcon("Restart AGS", theme.ViewRefreshIcon(), func() {
		restartags(currentUser)
	})
	btnSwaync := widget.NewButtonWithIcon("Restart SwayNC", theme.ViewRefreshIcon(), func() {
		restarswaync(currentUser)
		dialog.ShowInformation("Services", "SwayNC restarted!", w)
	})



	sidebarContent := container.NewVBox(
		title,
		configStatus,
		btnSelectConfig,
		widget.NewSeparator(),
		btnInstall,
		widget.NewLabelWithStyle("Color Schemes:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		pywalGrid,
		widget.NewSeparator(),
		widget.NewLabelWithStyle("Services:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		btnAgs,
		btnSwaync,
	)

	sidebar := container.NewPadded(sidebarContent)

	sidebarContainer := container.NewHBox(
		container.NewGridWrap(fyne.NewSize(320, 0), sidebar),
		widget.NewSeparator(),
	)

	w.SetContent(container.NewBorder(nil, nil, sidebarContainer, nil, img))
	w.ShowAndRun()
}
