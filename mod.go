package main

import "io"
import "archive/zip"
import "errors"
import "os/user"

var Public *zip.ReadCloser
var ActiveMod *Mod

type Mod struct {}

func init() {	
	var err error
	
	uid, _ := user.Current()
 
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
		for _, f := range Public.File {
			if f.Name == path {
				return f.Open()
			}
		}
	}
	return nil, errors.New("File not found!")
}
