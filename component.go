package main

import "github.com/icza/gowut/gwu"
import "strings"
import "sort"

func (e *Entity) ComponentEditor(component string, panel gwu.Panel) {
	Component := e.Get(component)

	if Component == nil {
		return
	}
	
	Keys := make([]string, 0)
	for k, _ := range Component {
		Keys = append(Keys, k)
	}
	sort.Strings(Keys)
	
	for _, key := range Keys {
		value := Component[key]
		child := key
		
		row := gwu.NewPanel()
		label := gwu.NewLabel(child)
		row.Add(label)
		
		//Special editors
		if strings.TrimSpace(value) == "true" || strings.TrimSpace(value) == "false" {
			
			checkbox := gwu.NewListBox([]string{"true", "false"})
			checkbox.SetSelected(0, value == "true")
			checkbox.SetSelected(1, value == "false")
			
			checkbox.AddEHandlerFunc(func(event gwu.Event) {
				e.Edit(component+"/"+child, checkbox.SelectedValue())
				SaveButton.SetText("Save")
				event.MarkDirty(SaveButton)
			}, gwu.ETypeChange)
			
			row.Add(checkbox)
			panel.Add(row)
			
		} else {
			if strings.TrimSpace(value) != "" {
				textbox := gwu.NewTextBox(value)
				textbox.AddSyncOnETypes(gwu.ETypeKeyUp)
				
				if strings.Contains(value, "\n") {
					textbox.SetRows(2)
					textbox.Style().SetWidth("400px")
					textbox.Style().SetHeightPx(16*strings.Count(value, "\n") + 10)
				}
				
				textbox.Style().SetWidth("300px")
				
				textbox.AddEHandlerFunc(func(event gwu.Event) {
					e.Edit(component+"/"+child, textbox.Text())
					SaveButton.SetText("Save")
					event.MarkDirty(SaveButton)
				}, gwu.ETypeChange, gwu.ETypeKeyUp)
				
				row.Add(textbox)
				panel.Add(row)
				
			
			} else { //This must have children elements.
				
				label.Style().SetFontSize("22px")
				panel.Add(row)
				
				child_panel := gwu.NewPanel()
				
				e.ComponentEditor(component+"/"+child, child_panel)
				
				child_panel.Style().SetPaddingLeft("20px")
				
				panel.Add(child_panel)
				
			}
		}
		
		//It's probably an image. So show it.
		//TODO work with user mods.
		if strings.Contains(value, ".png") {
			for _, path := range Public.Images() {
				if strings.Contains(path, value) {
					panel.Add(gwu.NewImage(path, Github+strings.Replace(path, ".cached.dds", "", 1)))
					break
				}
			}
		}
		
	}
}
