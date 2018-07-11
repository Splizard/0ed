package main

import "github.com/icza/gowut/gwu"

const Templates = "simulation/templates/"
const Portraits = "art/textures/ui/session/portraits/"
const Actors = "art/actors/"
const Meshes = "art/meshes/"
const Textures = "art/textures/skins/"

var Server gwu.Server

func main() {
	//GUI
	Server = gwu.NewServer("0ed", "localhost:8081")
	Server.SetText("0ed")
	
	Server.AddStaticDir("js", "./js")
	
	Server.AddWin(CreateStartupWindow())
	
	Server.Start("startup") // Also opens windows list in browser
}
