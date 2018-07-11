package main

import "os/user"
import "os"
import "io/ioutil"

import "github.com/icza/gowut/gwu"

const Templates = "simulation/templates/"
const Portraits = "art/textures/ui/session/portraits/"
const Actors = "art/actors/"
const Meshes = "art/meshes/"
const Textures = "art/textures/skins/"

var Mods string = "./mods"

func init() {	
	uid, _ := user.Current()
	
	if runtime.GOOS == "linux" {
		Mods = uid.HomeDir+"/.local/share/0ad/mods/0ed/"
	}
	
	if runtime.GOOS == "windows" {
		Mods = uid.HomeDir+"/Documents/My Games/0ad/mods/"
	}
	
	os.MkdirAll(Mods, 0755)
	
	ioutil.WriteFile(Mods+"mod.json", []byte(`{"name": "0ed","version": "0.1","label": "Sample Mod","description": "This is an example discription","dependencies": ["0ad>0.0.22"]} `), 0755)
}

var Server gwu.Server

func main() {
	//GUI
	Server = gwu.NewServer("0ed", "localhost:8081")
	Server.SetText("0ed")
	
	Server.AddStaticDir("js", "./js")
	
	Server.AddWin(CreateStartupWindow())
	
	Server.Start("startup") // Also opens windows list in browser
}
