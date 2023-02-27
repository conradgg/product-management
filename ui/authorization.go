package ui

import (
	"fmt"
	"product-management/user"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Authorization() {
	a := app.NewWithID("product-management")
	window := a.NewWindow("Product Management")
	window.Resize(fyne.NewSize(400, 250))
	window.CenterOnScreen()
	label := widget.NewLabel("Authorization")
	label.TextStyle = fyne.TextStyle{Bold: true}
	username := widget.NewEntry()
	username.SetPlaceHolder("Username")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")
	submit := widget.NewButton("Submit", func() {
		uName := strings.TrimSpace(username.Text)

		if uName == "" {

		} else {
			passwd := user.Encrypt(&password.Text)
			if user.Auth(&user.User{
				UserName: uName,
				Password: passwd,
			}) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	})
	reg := widget.NewButton("Create account?", func() {
		Register()
		window.Hide()
	})
	window.SetContent(container.NewVBox(
		label,
		username,
		password,
		submit,
		reg,
	))
	window.Show()
	a.Run()
}
