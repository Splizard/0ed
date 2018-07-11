package main

import "github.com/beevik/etree"

type Actor etree.Element 

func LoadActor(path string) (*Actor, error) {
	file, err := ActiveMod.Open(Actors+path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	tree := etree.NewDocument()
	if _, err := tree.ReadFrom(file); err != nil {
		return nil, err
	}
	
	return (*Actor)(tree.Root()), nil
}

func (a *Actor) Mesh() string {
	mesh := (*etree.Element)(a).FindElement("./group/variant/mesh")
	if mesh == nil {
		return ""
	}
	
	return Meshes+mesh.Text()
}

func (a *Actor) Props() map[string]string {
	props := (*etree.Element)(a).FindElement("./group/variant/props")
	if props == nil {
		return nil
	}
	
	result := make(map[string]string)
	for _, child := range props.ChildElements() {
		result[child.SelectAttr("actor").Value] = child.SelectAttr("attachpoint").Value
	}
	
	return result
}

func (a *Actor) Texture() string {
	texture := (*etree.Element)(a).FindElement("./group/variant/textures/texture[@name='baseTex']")
	if texture == nil {
		return ""
	}
	
	return Textures+texture.SelectAttr("file").Value
}
