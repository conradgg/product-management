package ui

import (
	"log"
	"os"

	"gioui.org/app"
)

func UI() {
	go func() {
		w := app.NewWindow()
		err := Authorization(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
