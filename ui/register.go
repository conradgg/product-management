package ui

import (
	"product-management/user"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Register() {
	regWindow := UI().NewWindow("Register")
	regWindow.Resize(fyne.NewSize(400, 500))
	label := widget.NewLabel("Register")
	label.TextStyle = fyne.TextStyle{Bold: true}
	username := widget.NewEntry()
	username.SetPlaceHolder("Username")
	firstname := widget.NewEntry()
	firstname.SetPlaceHolder("First Name")
	lastname := widget.NewEntry()
	lastname.SetPlaceHolder("Last Name")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")
	password1 := widget.NewPasswordEntry()
	password1.SetPlaceHolder("password")
	submit := widget.NewButton("Submit", func() {
		uName := strings.TrimSpace(username.Text)
		fName := strings.TrimSpace(firstname.Text)
		lName := strings.TrimSpace(lastname.Text)
		if uName == "" {

		} else if fName == "" {

		} else if lName == "" {

		} else if password.Text != password1.Text {

		} else {
			passwd := user.Encrypt(&password.Text)
			user.Register(&user.User{
				UserName:  uName,
				FirstName: fName,
				LastName:  lName,
				Password:  passwd,
			})
			Authorization()
		}
	})
	regWindow.SetContent(container.NewVBox(
		label,
		username,
		firstname,
		lastname,
		password,
		password1,
		submit,
	))
	regWindow.Show()
}
