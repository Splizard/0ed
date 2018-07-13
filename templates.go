package main

import "github.com/icza/gowut/gwu"

var HidePublic bool

func CreateStartupWindow() (Window gwu.Window) {
	Window = gwu.NewWindow("startup", "0ed")
	
	
	Window.AddHeadHTML(`
		<script src="js/jquery.min.js"></script>
		<link href="https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.6-rc.0/css/select2.min.css" rel="stylesheet" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.6-rc.0/js/select2.min.js"></script>
		
		<style>
			.select2 {
				position:absolute;
				top:calc(100vh / 2 - 180px);left:calc(100vw / 2 - 400px);
				font-family: "ubuntu", Times, serif;
				font-size: 12px;
			}
			
			.select2-results__option {
				font-size: 10px;
			}
			
			.select2-container--default .select2-results > .select2-results__options {
				max-height: 300px;
			}
		
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
			z-index: -99;
			}
			.model{
			position:absolute;
			top:calc(100vh / 2 - 200px);left:calc(100vw / 2 - 100px);
			background-color:black; width:400px; height:400px;
			color:white;font-family: "ubuntu", Times, serif;
			}
		</style>
    
	`)
	
	Select := gwu.NewHTML(`
	<script>
	
	
		$(document).ready(function() {
		
			$.fn.select2.defaults.set('matcher', function(params, data) {
				// If there are no search terms, return all of the data
				if ($.trim(params.term) === '') {
					return data;
				}

				// Do not display the item if there is no 'text' property
				if (typeof data.text === 'undefined') {
					return null;
				}

				var words = params.term.toUpperCase().split(" ");

				for (var i = 0; i < words.length; i++) {
					if (data.text.toUpperCase().indexOf(words[i]) < 0) {
					return null;
					}
				}

				return data;
			});
		
			$('.selectpicker').select2();
		});
	
	
</script>`)
	
	MoreWidgets(Window).DisableBackButton()
	
	var templates = Public.Templates()
	
	ListBox := gwu.NewListBox(append([]string{""}, templates...))
	ListBox.SetMulti(false)
		ListBox.Style().Set("position", "absolute").
		Set("top", "calc(100vh / 2 - 180px)").
		Set("left", "calc(100vw / 2 - 400px)").
		Set("z-index", "99").
		SetWidthPx(300).
		SetHeightPx(380)
	Window.Add(ListBox)
	ListBox.Style().AddClass("selectpicker")
	
	
	Window.Add(gwu.NewHTML(`
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
	
	var components = Public.Components()
	
	ComponentsListBox := gwu.NewListBox(components)

	//New ModManager() {
		
	
		var mod_manager = gwu.NewHorizontalPanel()
		listbox := gwu.NewListBox(append([]string{ActiveMod.Name}, OtherMods()...))
		Window.Add(gwu.NewLabel("Active Mod:"))
		mod_manager.Add(listbox)
		
		listbox.AddEHandlerFunc(func(e gwu.Event) {
			NewMod(listbox.SelectedValue())
		}, gwu.ETypeChange)
		
		var new_mod_name = gwu.NewTextBox("")
		
		mod_manager.Add(new_mod_name)
		
		var new_mod = gwu.NewButton("New")
		
		new_mod.AddEHandlerFunc(func(e gwu.Event) {
			if e.MouseBtn() == gwu.MouseBtnLeft {
				
				NewMod(new_mod_name.Text())
				listbox.SetValues(append([]string{ActiveMod.Name}, OtherMods()...))
				
				e.MarkDirty(listbox)
			}
		}, gwu.ETypeClick)
		
		mod_manager.Add(new_mod)

		CheckBox := gwu.NewCheckBox("Hide Public")
		CheckBox.SetState(HidePublic)
		mod_manager.Add(CheckBox)
		
		if HidePublic {
			ListBox.SetValues(ActiveMod.Templates())
			ComponentsListBox.SetValues(ActiveMod.Components())	
		}
		
		CheckBox.AddEHandlerFunc(func(e gwu.Event) {
			if e.MouseBtn() == gwu.MouseBtnLeft {
				
				HidePublic = CheckBox.State()
				
				if CheckBox.State() {
					e.MarkDirty(ListBox)
					e.MarkDirty(Select)
					ListBox.SetValues(ActiveMod.Templates())
					
					e.MarkDirty(ComponentsListBox)
					ComponentsListBox.SetValues(ActiveMod.Components())
				} else {
					e.MarkDirty(ListBox)
					e.MarkDirty(Select)
					ListBox.SetValues(Public.Templates())
					
					e.MarkDirty(ComponentsListBox)
					ComponentsListBox.SetValues(Public.Components())
				}
			}
		}, gwu.ETypeClick)
		
		Window.Add(mod_manager)
	// }
	
	
	Window.Add(gwu.NewLabel("Component Editor"))
	
	
	Button = gwu.NewButton("Edit Component")
			
	Button.AddEHandlerFunc(func(e gwu.Event) {
		if e.MouseBtn() == gwu.MouseBtnLeft {
			Server.RemoveWin(Window)
			Server.AddWin(CreateComponentEditorWindow(ComponentsListBox.SelectedValue()))
			e.ReloadWin("ceditor")
		}
	}, gwu.ETypeClick)
	Window.Add(Button)
	
	
	ComponentsListBox.SetMulti(true)
	ComponentsListBox.Style().SetWidthPx(200)
	ComponentsListBox.Style().SetHeightPx(800)
	Window.Add(ComponentsListBox)
	
	Window.Add(Select)
	
	return
}
