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

func (e *Entity) Edit(path string, value string) {
	element := e.Components.FindElement("./"+path)
	
	if element == nil {
		reader := bufio.NewReader(strings.NewReader(path))
		last := ""
		for i:=0; i < strings.Count(path, "/"); i++ {
			dir, _ := reader.ReadString('/')
			dir = dir[:len(dir)-1]

			if element = e.Components.FindElement("./"+dir); element == nil {
				if last == "" {
					element = e.Components.CreateElement(dir)
				} else {
					element = e.Components.FindElement("./"+last).CreateElement(dir)
				}
			} 
			last = dir
		}
		
		element.CreateElement(filepath.Base(path)).SetText(value)
		
	} else {
		element.SetText(value)
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
			
			if child.SelectAttr("disable") == nil {
				result[child.Tag] = child.Text()
			} else {
				delete(result, child.Tag)
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
