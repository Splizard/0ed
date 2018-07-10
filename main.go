package main

import "net/http"
import "os/user"
import "os"
import "io/ioutil"

import "github.com/icza/gowut/gwu"

const Public = "../0ad/binaries/data/mods/public/"
const Templates = "simulation/templates/"
const Portraits = "art/textures/ui/session/portraits/"
const Actors = "art/actors/"
const Meshes = "art/meshes/"
const Textures = "art/textures/skins/"

var Mod string

func init() {
	uid, _ := user.Current()
	Mod = uid.HomeDir+"/.local/share/0ad/mods/0ed/"
	
	os.MkdirAll(Mod, 0755)
	
	ioutil.WriteFile(Mod+"mod.json", []byte(`{"name": "0ed","version": "0.1","label": "Sample Mod","description": "This is an example discription","dependencies": ["0ad>0.0.22"]} `), 0755)
}

var Server gwu.Server

func main() {
	
	go http.ListenAndServe("localhost:8080", http.FileServer(http.Dir(Public)))
	
	//GUI
	Server = gwu.NewServer("0ed", "localhost:8081")
	Server.SetText("0ed")
	
	Server.AddStaticDir("js", "./js")
	Server.AddStaticDir("test", "./test")
	Server.AddStaticDir("art", Public+"/art")
	
	Server.AddWin(CreateStartupWindow())
	
	Server.Start("startup") // Also opens windows list in browser
}
