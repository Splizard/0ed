package main

import "github.com/beevik/etree"

type Actor etree.Element 

func LoadActor(path string) (*Actor, error) {
	tree := etree.NewDocument()
	if err := tree.ReadFromFile(path); err != nil {
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

func (a *Actor) Texture() string {
	texture := (*etree.Element)(a).FindElement("./group/variant/textures/texture[@name='baseTex']")
	if texture == nil {
		return ""
	}
	
	return Textures+texture.SelectAttr("file").Value
}
