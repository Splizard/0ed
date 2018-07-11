package main

import "io"
import "archive/zip"
import "errors"
import "os/user"
import "os"
import "io/ioutil"
import "runtime"

var Public *zip.ReadCloser
var ActiveMod *Mod

type Mod struct {}

var Mods string = "./mods"

func init() {	
	uid, _ := user.Current()
	
	if runtime.GOOS == "linux" {
		Mods = uid.HomeDir+"/.local/share/0ad/mods/0ed/"
	}
	
	if runtime.GOOS == "windows" {
		Mods = uid.HomeDir+"/Documents/My Games/0ad/mods/0ed/"
	}
	
	os.MkdirAll(Mods, 0755)
	
	ioutil.WriteFile(Mods+"mod.json", []byte(`{"name": "0ed","version": "0.1","label": "Sample Mod","description": "This is an example discription","dependencies": ["0ad>0.0.22"]} `), 0755)
	
	var err error
	//Try using 0ad's public.zip file.
	// Open a zip archive for reading.
	Public, err = zip.OpenReader("/usr/share/0ad/data/mods/public/public.zip")

		if err != nil {
		
	Public, err = zip.OpenReader(uid.HomeDir+"/AppData/Local/0 A.D. alpha/")
	
		if err != nil {
			println("Could not locate 0ad data file!")
		}

	}
}

func (m *Mod) Open(path string) (io.ReadCloser, error) {
	if m == nil {
		
		if file, err := os.Open(Mods+path); err == nil {
			return file, err
		}
		
		for _, f := range Public.File {
			if f.Name == path {
				return f.Open()
			}
		}
	}
	return nil, errors.New("File not found!")
}
