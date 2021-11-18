package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var (
	selections = []string{
		"Alpha",
		"Bravo",
	}

	contents = map[string][]string{"Alpha": {"Alpha one", "Alpha two", "Alpha three"}, "Bravo": {"Bravo one", "Bravo two", "Bravo three"}}
)

func selectorWidget(binder binding.String) *widget.Select {
	return widget.NewSelect(selections, func(value string) {
		binder.Set(value)
	})
}

func listenerCallback(binder binding.String, contentPanel *fyne.Container) func() {
	return func() {
		selection, err := binder.Get()
		for _, obj := range contentPanel.Objects {
			contentPanel.Remove(obj)
		}
		if err == nil {
			for _, content := range contents[selection] {
				label := widget.NewLabel(content)
				contentPanel.Add(label)
			}
		}
	}
}

func main() {

	// Initialise app
	myApp := app.New()
	myWindow := myApp.NewWindow(("Selector example"))

	// Data binder
	binder := binding.NewString()

	// Selector panel
	selWidget := selectorWidget(binder)
	selLabelWidget := widget.NewLabel("Please select: ")
	selPanel := container.NewHBox()
	selPanel.Add(selLabelWidget)
	selPanel.Add(selWidget)

	// Content panel
	contentPanel := container.NewVBox()
	binder.AddListener(binding.NewDataListener(listenerCallback(binder, contentPanel)))

	// Main panel
	mainPanel := container.NewVBox()
	mainPanel.Add(selPanel)
	mainPanel.Add(contentPanel)

	// Windows
	myWindow.SetContent(mainPanel)
	myWindow.Resize(fyne.NewSize(700, 400))
	myWindow.ShowAndRun()
}
