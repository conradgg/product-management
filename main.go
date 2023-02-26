package main

import "fyne.io/fyne/v2/app"

func main() {
	a := app.New()
	w := a.NewWindow("Product Management")
	w.ShowAndRun()
}
