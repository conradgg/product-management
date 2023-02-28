package ui

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type C = layout.Context
type D = layout.Dimensions

func Authorization() {
	go func() {
		w := app.NewWindow(
			app.Size(400, 250),
			app.Title("Product Management"),
			app.MinSize(350, 250),
		)
		if err := authWindow(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func authWindow(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops
	var startButton widget.Clickable
	var usernameInput widget.Editor
	var passwordInput widget.Editor
	margins := layout.UniformInset(15)
	border := widget.Border{
		Color:        color.NRGBA{R: 0, G: 0, B: 0, A: 200},
		CornerRadius: 7,
		Width:        0.5,
	}

	for e := range w.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceSides,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx C) D {
						title := material.H3(th, "Authorization")
						title.Alignment = text.Middle
						return title.Layout(gtx)
					},
				),
				layout.Rigid(
					func(gtx C) D {
						username := material.Editor(th, &usernameInput, " Username")
						usernameInput.SingleLine = true
						return margins.Layout(gtx,
							func(gtx C) D {
								return border.Layout(gtx, username.Layout)
							},
						)
					},
				),
				layout.Rigid(layout.Spacer{Height: unit.Dp(25)}.Layout),
				layout.Rigid(
					func(gtx C) D {
						password := material.Editor(th, &passwordInput, "Password")
						passwordInput.SingleLine = true
						return password.Layout(gtx)
					},
				),
				layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
				layout.Rigid(
					func(gtx C) D {
						btn := material.Button(th, &startButton, "Submit")
						return btn.Layout(gtx)
					},
				),
			)
			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}

	}
	return nil
}
