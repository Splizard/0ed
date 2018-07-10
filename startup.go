package main

import "github.com/icza/gowut/gwu"

func CreateTemplatesWindow() (Window gwu.Window) {
	var Label gwu.Label
	var Button gwu.Button
	
	Window = gwu.NewWindow("startup2", "Startup2")
	Label = gwu.NewLabel("Choose to create a new actor, or edit an existing one")
		Label.Style().Set("position", "absolute").
			Set("top", "calc(100vh / 2 - 70px)").
			Set("left", "calc(100vw / 2 - 190px)")
	Window.Add(Label)
	
	Button = gwu.NewButton("New")
		Button.Style().Set("position", "absolute").
			Set("top", "calc(100vh / 2 - 25px)").
			Set("left", "calc(100vw / 2 - 200px)").
			SetWidthPx(200).
			SetHeightPx(50).
			SetBackground("red").
			Set("border-radius", "8px").
			Set("color", "white")
			
	Button.AddEHandlerFunc(func(e gwu.Event) {
		if e.MouseBtn() == gwu.MouseBtnLeft {
			Server.RemoveWin(Window)
			Server.AddWin(CreateEditorWindow(""))
			e.ReloadWin("editor")
		}
	}, gwu.ETypeClick)
	Window.Add(Button)
		
	Button = gwu.NewButton("Edit")
		Button.Style().Set("position", "absolute").
			Set("top", "calc(100vh / 2 - 25px)").
			Set("left", "calc(100vw / 2)").
			SetWidthPx(200).
			SetHeightPx(50).
			SetBackground("green").
			Set("border-radius", "8px").
			Set("color", "white")
	Button.AddEHandlerFunc(func(e gwu.Event) {
		if e.MouseBtn() == gwu.MouseBtnLeft {
			Server.RemoveWin(Window)
			Server.AddWin(CreateTemplatesWindow())
			e.ReloadWin("templates")
		}
	}, gwu.ETypeClick)
	Window.Add(Button)
	
	return
}
