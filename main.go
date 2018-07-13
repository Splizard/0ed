package main

import "github.com/icza/gowut/gwu"

const Templates = "simulation/templates/"
const Portraits = "art/textures/ui/session/portraits/"
const Actors = "art/actors/"
const Meshes = "art/meshes/"
const Sounds = "audio/"
const Textures = "art/textures/skins/"
const Components = "simulation/components/"
const Github = "https://raw.githubusercontent.com/0ad/0ad/master/binaries/data/mods/public/"

var Server gwu.Server

func main() {
	ActiveMod.Load()
	
	//GUI
	Server = gwu.NewServer("0ed", "localhost:8081")
	Server.SetText("0ed")
	
	Server.AddStaticDir("js", "./js")
	
	Server.AddWin(CreateStartupWindow())
	
	Server.Start("startup") // Also opens windows list in browser
}
