package main

import "github.com/icza/gowut/gwu"
import "path/filepath"
import "os"
import "strings"

func CreateStartupWindow() (Window gwu.Window) {
	Window = gwu.NewWindow("startup", "0ed")
	
	var templates []string
	
	filepath.Walk(Public+Templates, func(path string, info os.FileInfo, err error) error {
		name := strings.Split(path, Public+Templates)[1]
		if len(name) > 4 {
		templates = append(templates, name[:len(name)-4])
		}
		return nil
	})
	
	ListBox := gwu.NewListBox(templates)
	ListBox.SetMulti(true)
		ListBox.Style().Set("position", "absolute").
		Set("top", "calc(100vh / 2 - 180px)").
		Set("left", "calc(100vw / 2 - 400px)").
		Set("z-index", "99").
		SetWidthPx(300).
		SetHeightPx(380)
	Window.Add(ListBox)
	
	Window.Add(gwu.NewHTML(`<style>
    @font-face {
      font-family: 'ubuntu'; /*a name to be used later*/
      src: url('test/Ubuntu-R.ttf'); /*URL to font*/
    }
    .accept{
      position:absolute;
      top:calc(100vh / 2 + 150px);left:calc(100vw / 2 - 100px);
      background-color:green; width:400px; height:50px;
      border-radius: 8px; color:white; font-family: "ubuntu", Times, serif;
    }
    .accept:hover{
      background-color: #4CAF50; /* Green */
    }
    .actor{
      position:absolute;
      top:calc(100vh / 2 - 200px);left:calc(100vw / 2 - 400px);
      background-color:grey; width:400px; height:400px;
      font-family: "ubuntu", Times, serif;
    }
    .model{
      position:absolute;
      top:calc(100vh / 2 - 200px);left:calc(100vw / 2 - 100px);
      background-color:black; width:400px; height:400px;
      color:white;font-family: "ubuntu", Times, serif;
    }
    </style>
    
    <div class="actor">
         <center>Select an actor</center>
    </div>
    <div class="model">
         <center>Model</center>
    </div>

  `))
	
		Button := gwu.NewButton("Accept")
	Button.Style().SetClass("accept")
			
	Button.AddEHandlerFunc(func(e gwu.Event) {
		if e.MouseBtn() == gwu.MouseBtnLeft {
			Server.RemoveWin(Window)
			Server.AddWin(CreateEditorWindow(ListBox.SelectedValue()))
			e.ReloadWin("editor")
		}
	}, gwu.ETypeClick)
	Window.Add(Button)
	
	listbox := gwu.NewListBox([]string{"0ed"})
	Window.Add(gwu.NewLabel("Active Mod:"))
	Window.Add(listbox)
	
	return
}
