package main

import "github.com/icza/gowut/gwu"
import "fmt"
import "os"
import "os/exec"
import "sort"

var SaveButton gwu.Button

var EditorRefresh func(e gwu.Event)

func CreateEditorWindow(template string) (Window gwu.Window) {
	entity, err := LoadEntity(template)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	
	actor, err := LoadActor(entity.Actor())
	if err != nil {
		fmt.Println(err.Error())
	}
	
	//MainWindow
	Window = gwu.NewWindow("editor", "Entity Editor")
	MoreWidgets(Window).DisableBackButton()
	
	Window.AddHeadHTML(`
	<style>
	.gwu-Window {
		position:absolute;
		top:100px;left:400px;
		font-family: "ubuntu", Times, serif;
	}
	
    @font-face {
      font-family: 'ubuntu'; /*a name to be used later*/
      src: url('./Ubuntu-R.ttf'); /*URL to font*/
    }
    .add{
      position:fixed;
      top:60px;left:440px;
      background-color:blue; width:30px; height:30px;
      border-radius: 4px; color:white; font-family: "ubuntu", Times, serif;
    }
    .add:hover{
      background-color: #575edb; /* Blue */
    }
    .back{
      position:fixed;
      top:60px;left:405px;
      background-color:blue; width:30px; height:30px;
      border-radius: 4px; color:white; font-family: "ubuntu", Times, serif;
    }
    .back:hover{
      background-color: #575edb; /* Blue */
    }
    .save{
      position:fixed;
      top:0px;left:500px;
      background-color:green; width:200px; height:50px;
      border-radius: 8px; color:white; font-family: "ubuntu", Times, serif;
    }
    .save:hover{
      background-color: #4CAF50; /* Green */
    }

    .run{
      position:fixed;
      top:0px;left:700px;
      background-color:red; width:200px; height:50px;
      border-radius: 8px; color:white; font-family: "ubuntu", Times, serif;
    }
    .run:hover{
      background-color: #f76767; /* Red */
    }
    .home{
      position:fixed;
      top:0px;left:400px;
      background-color:blue; width:100px; height:50px;
      border-radius: 8px; color:white; font-family: "ubuntu", Times, serif;
    }
    .home:hover{
      background-color: #575edb; /* Blue */
    }
    .modelEditor{
      position:fixed;
      top:0px;left:0px;
      background-color:grey; width:400px; height:calc(100vh - 400px);
      font-family: "ubuntu", Times, serif;
    }
    .edit{
      position:fixed;
      top:50px;left:400px;
      background-color:silver; width:calc(100vw - 400px); height:calc(100vh - 50px);
      font-family: "ubuntu", Times, serif;
	  z-index: -99;
    }
    </style>
	
	`)
	
	//Main Editor, content will change when a compnent is clicked.
	editor := gwu.NewPanel()
	editor.SetHAlign(gwu.HALeft)
	
	//Components.
	components := gwu.NewNaturalPanel()
	
	Map := entity.Map()
	Keys := make([]string, 0)
	for k, _ := range Map {
		Keys = append(Keys, k)
	}
	sort.Strings(Keys)
	
	for _, key := range Keys {
		component := key
		
		button := gwu.NewButton(component)
		
		button.Style().Set("float", "left")
		components.Add(button)
		
		if !entity.State(component) {
			button.SetText(component+" ❌")
		}
		
		var refresh = func(e gwu.Event) {
			//if e.MouseBtn() == gwu.MouseBtnLeft {
				editor.Clear()
				editor.Style().SetPadding("50px")
				label := gwu.NewLabel(component)
				label.Style().SetFontSize("32px")
				label.Style().SetPaddingBottom("30px")
				
				row := gwu.NewHorizontalPanel()
				
				row.Add(label)
				
				
				DisableButton := gwu.NewButton("✓")
				DisableButton.Style().SetMarginLeft("25px")
				
				if entity.State(component) {
					DisableButton.SetText("✓")
				} else {
					DisableButton.SetText("❌")
				}
				
				DisableButton.AddEHandlerFunc(func(e gwu.Event) {
					e.MarkDirty(DisableButton)
					e.MarkDirty(button)
					
					if DisableButton.Text() == "✓" {
						DisableButton.SetText("❌")
						button.SetText(component+" ❌")
						entity.Disable(component)
					} else {
						DisableButton.SetText("✓")
						button.SetText(component)
						entity.Enable(component)
					}
				}, gwu.ETypeClick)
				
				row.Add(DisableButton)
				
				editor.Add(row)
				editor.SetHAlign(gwu.HALeft)
				
				
				
				//Generate the component editor for this component.
				entity.ComponentEditor(component, editor)
				
				e.MarkDirty(editor)
			//}
		}
			
		button.AddEHandlerFunc(func(e gwu.Event) {
			EditorRefresh = refresh
			refresh(e)
		}, gwu.ETypeClick)
	}
	Window.Add(components)
	Window.Add(editor)
	
	
	
	ModelViewer := NewModelViewer(actor)
	Window.Add(ModelViewer)
	
	Window.Add(gwu.NewHTML(`

    <div class="modelEditor">
         <p> </p>
         <center>`+template+`</center>
    </div>
    
    <button type="button" class="back">←</button>
    <button type="button" class="add">+</button>
    <form method="post" action="startup">
      <button type="submit" class="home">HOME</button>
    </form>
    
    <div class="edit">
         <center>Component Editor</center>
    </div>
	`))
	
	/*img := gwu.NewImage("Icon", "http://localhost:8080/"+Portraits+ent.Identity.Icon)
	img.Style().SetSizePx(100, 100)
	Window.Add(img)
	
	Window.Add(gwu.NewLabel(fmt.Sprint(ent.Identity.Tooltip)))
	
	Window.Add(gwu.NewLabel(fmt.Sprint(ent.Cost)))*/
	
	Button := gwu.NewButton("Home")
	Button.Style().SetClass("home")
			
	Button.AddEHandlerFunc(func(e gwu.Event) {
		if e.MouseBtn() == gwu.MouseBtnLeft {
			Server.RemoveWin(Window)
			Server.AddWin(CreateStartupWindow())
			e.ReloadWin("startup")
		}
	}, gwu.ETypeClick)
	Window.Add(Button)
	
	SaveButton = gwu.NewButton("Save")
	SaveButton.Style().SetClass("save")
			
	SaveButton.AddEHandlerFunc(func(e gwu.Event) {
		if e.MouseBtn() == gwu.MouseBtnLeft {
			SaveButton.SetText("Saved ✓")
			e.MarkDirty(SaveButton)
			
			entity.WriteToFile(Mods+Templates+template)
		}
	}, gwu.ETypeClick)
	Window.Add(SaveButton)
	
	Button = gwu.NewButton("Run")
	Button.Style().SetClass("run")
			
	Button.AddEHandlerFunc(func(e gwu.Event) {
		if e.MouseBtn() == gwu.MouseBtnLeft {
			fmt.Println(exec.Command("pyrogenesis", "-mod=public",  "-mod=0ed", "-autostart=skirmishes/Acropolis Bay (2)", "-autostart-civ=1:athen").Start())

		}
	}, gwu.ETypeClick)
	Window.Add(Button)
	
	return
}
