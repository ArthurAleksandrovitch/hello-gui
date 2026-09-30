package main

import (
	"fmt"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var version = "dev"

func main() {
	a := app.New()
	w := a.NewWindow("Hello GUI")

	versionLabel := widget.NewLabel(
		fmt.Sprintf("Version: %s", version),
	)

	systemLabel := widget.NewLabel(
		fmt.Sprintf("OS: %s | Arch: %s", runtime.GOOS, runtime.GOARCH),
	)

	greetingLabel := widget.NewLabel(
		Greeting("Fyne"),
	)

	sumLabel := widget.NewLabel(
		fmt.Sprintf("Sum 1..10 = %d", SumRange(1, 10)),
	)

	greetButton := widget.NewButton("Greet", func() {
		greetingLabel.SetText(Greeting("Fyne"))
	})

	quitButton := widget.NewButton("Quit", func() {
		a.Quit()
	})

	content := container.NewVBox(
		widget.NewLabel("Hello from Go GUI! 🐹"),
		versionLabel,
		systemLabel,
		greetingLabel,
		sumLabel,
		greetButton,
		quitButton,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(500, 350))
	w.ShowAndRun()
}
