package main

import "github.com/beevik/etree"
import "strings"

var ComponentPalette *etree.Document

func init() {
	ComponentPalette = etree.NewDocument()
	if _, err := ComponentPalette.ReadFrom(strings.NewReader(ComponentPaletteRaw)); err != nil {
		println(err.Error())
	}
}
