package main

import "github.com/icza/gowut/gwu"
import "fmt"
import "os"
import "os/exec"
import "path/filepath"

import "io/ioutil"


func CreateComponentEditorWindow(path string) (Window gwu.Window) {
	//MainWindow
	Window = gwu.NewWindow("ceditor", "Component Editor")
	
	file, err := ActiveMod.Open(Components+path+".js")
	if err != nil {
		Window.Add(gwu.NewLabel(fmt.Sprint(err)))
		return
	}
	defer file.Close()
	
	data, err := ioutil.ReadAll(file)
	if err != nil {
		Window.Add(gwu.NewLabel(fmt.Sprint(err)))
		return
	}
	
	Code := gwu.NewTextBox(string(data))
	
	Button := gwu.NewButton("Home")
	//Button.Style().SetClass("home")
			
	Button.AddEHandlerFunc(func(e gwu.Event) {
		if e.MouseBtn() == gwu.MouseBtnLeft {
			Server.RemoveWin(Window)
			Server.AddWin(CreateStartupWindow())
			e.ReloadWin("startup")
		}
	}, gwu.ETypeClick)
	Window.Add(Button)
	
	Button = gwu.NewButton("Run")
	Button.Style().SetClass("run")
			
	Button.AddEHandlerFunc(func(e gwu.Event) {
		if e.MouseBtn() == gwu.MouseBtnLeft {
			fmt.Println(exec.Command("pyrogenesis", "-mod=public",  "-mod=0ed", "-autostart=skirmishes/Acropolis Bay (2)", "-autostart-civ=1:athen").Start())

		}
	}, gwu.ETypeClick)
	Window.Add(Button)
	
	SaveButton := gwu.NewButton("Save")
	SaveButton.Style().SetClass("save")
			
	SaveButton.AddEHandlerFunc(func(e gwu.Event) {
		if e.MouseBtn() == gwu.MouseBtnLeft {
			SaveButton.SetText("Saved ✓")
			e.MarkDirty(SaveButton)
			
			os.MkdirAll(filepath.Dir(Mods+Components+path+".js"), 0755)
			ioutil.WriteFile(Mods+Components+path+".js", []byte(Code.Text()), 0755)
		}
	}, gwu.ETypeClick)
	Window.Add(SaveButton)
	
	
	Code.Style().SetWidth("99vw")
	Code.Style().SetHeight("90vh")
	Code.SetRows(2)
	
	Code.AddEHandlerFunc(func(event gwu.Event) {
		SaveButton.SetText("Save")
		event.MarkDirty(SaveButton)
	}, gwu.ETypeKeyUp)

	Window.Add(Code)
	

	

	
	return
}
