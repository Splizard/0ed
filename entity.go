package main

import "github.com/beevik/etree"
import "fmt"
import "path/filepath"
import "os"
import "strings"
import "bufio"

type Entity struct {
	tree *etree.Document
	
	Components *etree.Element
	
	Parent *Entity
}

func (e *Entity) WriteToFile(path string) {
	os.MkdirAll(filepath.Dir(path), 0755)
	if err := e.tree.WriteToFile(path+".xml"); err != nil {
		fmt.Println(err)
	}
}

//Enabled or Disabled components.
func (e *Entity) State(path string) bool {
	var attr string
	if strings.Contains(path, ".") {
		attr = strings.Split(path, ".")[1]
		path = strings.Split(path, ".")[0]
	}
	
	if attr != "" {
		println("cannot get the state of attributes!")
		return true
	}
	
	element := e.Components.FindElement("./"+path)
	
	if element == nil {
		
		//TODO
		
	} else {
		return element.SelectAttr("disable") == nil
	}
	
	return true
}

func (e *Entity) Disable(path string) {
	
	var attr string
	if strings.Contains(path, ".") {
		attr = strings.Split(path, ".")[1]
		path = strings.Split(path, ".")[0]
	}
	
	if attr != "" {
		println("cannot disable attributes!")
		return
	}
	
	element := e.Components.FindElement("./"+path)
	
	if element == nil {
				
		if !strings.Contains(path, "/") {
			e.Components.CreateElement(path).CreateAttr("disable", "")
			return
		}
		
		reader := bufio.NewReader(strings.NewReader(path))
		last := ""
		for i:=0; i < strings.Count(path, "/"); i++ {
			dir, _ := reader.ReadString('/')
			if last != "" {
				dir = last+"/"+dir[:len(dir)-1]
			} else {
				dir = dir[:len(dir)-1]
			}

			if element = e.Components.FindElement("./"+dir); element == nil {
				if last == "" {
					element = e.Components.CreateElement(filepath.Base(dir))
				} else {
					element = e.Components.FindElement("./"+last).CreateElement(filepath.Base(dir))
				}
			} 
			last = dir
		}
		
		element.CreateElement(filepath.Base(path)).CreateAttr("disable", "")
	} else {
		element.CreateAttr("disable", "")
	}
}

func (e *Entity) Enable(path string) {
	var attr string
	if strings.Contains(path, ".") {
		attr = strings.Split(path, ".")[1]
		path = strings.Split(path, ".")[0]
	}
	
	if attr != "" {
		println("cannot disable attributes!")
		return
	}
	
	element := e.Components.FindElement("./"+path)
	
	if element == nil {
		//Uh oh!
		println("Cannot enable non-existant element!")
	} else {
		if len(element.Child) == 0 && len(element.Attr) == 1 && e.Parent != nil && e.Parent.Component(filepath.Base(path)) != nil {
			element.Parent().RemoveChild(element)
		} else {
			element.RemoveAttr("disable")
		}
	}
}

func (e *Entity) Component(name string) *etree.Element {
	if element := e.Components.SelectElement(name); element != nil {
		return element
	}
	
	if e.Parent == nil {
		return nil
	}
	
	return e.Parent.Component(name)
}
	
func (e *Entity) Reset(path string) {
	
	var attr string
	if strings.Contains(path, ".") {
		attr = strings.Split(path, ".")[1]
		path = strings.Split(path, ".")[0]
	}
	
	element := e.Components.FindElement("./"+filepath.Dir(path))

	if element == nil {

	} else {
		if attr != "" {
			if element.SelectElement(filepath.Base(path)).SelectAttr(attr) != nil {
				element.SelectElement(filepath.Base(path)).RemoveAttr(attr)
			}
		} else {
			element.RemoveChild(element.SelectElement(filepath.Base(path)))
		}
	}
}

func (e *Entity) Edit(path string, value string) {
	
	var attr string
	if strings.Contains(path, ".") {
		attr = strings.Split(path, ".")[1]
		path = strings.Split(path, ".")[0]
	}
	
	element := e.Components.FindElement("./"+path)

	if element == nil {
		reader := bufio.NewReader(strings.NewReader(path))
		last := ""
		for i:=0; i < strings.Count(path, "/"); i++ {
			dir, _ := reader.ReadString('/')
			if last != "" {
				dir = last+"/"+dir[:len(dir)-1]
			} else {
				dir = dir[:len(dir)-1]
			}

			if element = e.Components.FindElement("./"+dir); element == nil {
				if last == "" {
					element = e.Components.CreateElement(filepath.Base(dir))
				} else {
					element = e.Components.FindElement("./"+last).CreateElement(filepath.Base(dir))
				}
			} 
			last = dir
		}
		
		if attr != "" {
			element.CreateElement(filepath.Base(path)).CreateAttr(attr, value)
		} else {
			element.CreateElement(filepath.Base(path)).SetText(value)
		}
		
	} else {
		if attr != "" {
			element.SelectAttr(attr).Value = value
		} else {
			element.SetText(value)
		}
	}
}

func (e *Entity) Get(component string) (result map[string]string) {
	if e.Parent != nil {
		result = e.Parent.Get(component)
	}
	if result == nil {
		result = make(map[string]string)
	}
	if e.Components.FindElement(component) != nil {
		for _, child := range e.Components.FindElement(component).ChildElements() {
			name := child.Tag
			if child.SelectAttr("disable") == nil {
				
				
			
				if len(child.Attr) > 0 {
					for _, attr := range child.Attr {
						if attr.Key == "disable" || attr.Key == "replace" {
							continue
						}
						result[name+"."+attr.Key] = attr.Value
					}
				}
				
				result[name] = child.Text()
			} else {
				delete(result, child.Tag)
				result[name] = child.Text()+"❌"
			}
		}
	}
	return result
}

func (e *Entity) Map() (components map[string]bool) {
	components = make(map[string]bool)
	
	for _, component := range e.Components.ChildElements() {
		if component.SelectAttr("disable") == nil {
			components[component.Tag] = true
		}
	}
	
	if e.Parent != nil {
		for component := range e.Parent.Map() {
			components[component] = true
		}
	}
	
	return
}

func (e *Entity) Name() string {
	
	if identity := e.Components.SelectElement("Identity"); identity != nil {
		if name := identity.SelectElement("SpecificName"); name != nil {
			
			if generic := identity.SelectElement("GenericName"); generic != nil {
				return generic.Text()+" ("+name.Text()+")"
			}
			
			return name.Text()
		}
		
		if name := identity.SelectElement("GenericName"); name != nil {
			return name.Text()
		}
	}
	
	if e.Parent != nil {
		return e.Parent.Name()
	}
	
	return ""
}

func (e *Entity) Actor() string {
	VisualActor := e.Components.SelectElement("VisualActor")
	if VisualActor == nil {
		
		if e.Parent != nil {
			return e.Parent.Actor()
		}
		
		return ""
	}
	
	actor := VisualActor.SelectElement("Actor")
	if actor == nil {
		
		if e.Parent != nil {
			return e.Parent.Actor()
		}
		
		return ""
	}
	return actor.Text()
}

func LoadEntity(template string) (e *Entity, err error) {
	file, err := ActiveMod.Open(Templates+template+".xml")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	tree := etree.NewDocument()
	if _, err = tree.ReadFrom(file); err != nil {
		return nil, err
	}
	
	e = new(Entity)
	e.tree = tree
	
	e.Components = tree.Root()
	if tree.Root().SelectAttr("parent") != nil {
		e.Parent, err = LoadEntity(tree.Root().SelectAttr("parent").Value)
	}
	return
}
