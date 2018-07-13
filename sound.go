package main

import "github.com/beevik/etree"

type Sound etree.Element 

func LoadSound(path string) (*Sound, error) {
	file, err := ActiveMod.Open(Sounds+path)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	defer file.Close()
	
	tree := etree.NewDocument()
	if _, err := tree.ReadFrom(file); err != nil {
		return nil, err
	}
	
	return (*Sound)(tree.Root()), nil
}

func (s *Sound) Random() string {
	path := (*etree.Element)(s).FindElement("./Path")
	if path != nil {
		element := (*etree.Element)(s).FindElement("./Sound")
		if element != nil {
			return (path.Text()+"/"+element.Text())
		}
	}
	
	return ""
}
