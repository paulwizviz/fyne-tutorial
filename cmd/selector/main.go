package main

import (
	"fmt"

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

	contentOne = []string{"Alpha one", "Alpha two", "Alpha three"}
	contentTwo = []string{"Bravo one", "Bravo two", "Bravo three"}
)

func selectorContainer(binder binding.StringList) *fyne.Container {

	sel := widget.NewSelect(selections, func(value string) {
		if value == selections[0] {
			binder.Set(contentOne)
		} else {
			binder.Set(contentTwo)
		}
	})

	cont := container.NewHBox()
	cont.Resize(fyne.NewSize(200, 100))
	cont.Add(sel)
	return cont
}

func content(binder binding.StringList) func() {
	return func() {
		content, err := binder.Get()
		fmt.Println(content, err)
	}
}

func main() {

	// Initialise app
	myApp := app.New()
	myWindow := myApp.NewWindow(("Selector example"))

	// Data binder
	binder := binding.NewStringList()

	// Content
	contentPanel := container.NewVBox()

	// Listener
	listener := binding.NewDataListener(content(binder))
	binder.AddListener(listener)

	// Main Panel
	mainPanel := container.NewVBox()
	mainPanel.Add(selectorContainer(binder))
	mainPanel.Add(contentPanel)

	// Windows

	myWindow.SetContent(mainPanel)
	myWindow.Resize(fyne.NewSize(700, 400))
	myWindow.ShowAndRun()
}
