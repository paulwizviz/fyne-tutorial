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

func selectorWidget(binder binding.StringList) *widget.Select {

	sel := widget.NewSelect(selections, func(value string) {
		if value == selections[0] {
			binder.Set(contentOne)
		} else {
			binder.Set(contentTwo)
		}
	})

	return sel
}

func listenerCallback(binder binding.StringList, labelList *fyne.Container) func() {
	return func() {
		dataItems, err := binder.Get()
		fmt.Println(dataItems)
		if err == nil {
			for _, dataItem := range dataItems {
				label := widget.NewLabel(dataItem)
				labelList.Add(label)
			}
		}
	}
}

func main() {

	// Initialise app
	myApp := app.New()
	myWindow := myApp.NewWindow(("Selector example"))

	// Data binder
	binder := binding.NewStringList()

	// Selector panel
	selWidget := selectorWidget(binder)
	selLabelWidget := widget.NewLabel("Please select: ")
	selPanel := container.NewHBox()
	selPanel.Add(selLabelWidget)
	selPanel.Add(selWidget)

	// Content panel
	contentPanel := container.NewVBox()
	cb := listenerCallback(binder, contentPanel)
	listener := binding.NewDataListener(cb)
	binder.AddListener(listener)

	// Main panel
	mainPanel := container.NewVBox()
	mainPanel.Add(selPanel)
	mainPanel.Add(contentPanel)

	// Windows
	myWindow.SetContent(mainPanel)
	myWindow.Resize(fyne.NewSize(700, 400))
	myWindow.ShowAndRun()
}
