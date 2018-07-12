package main

import "io"
import "archive/zip"
import "errors"
import "os/user"
import "os"
import "io/ioutil"
import "runtime"
import "strings"
import "sort"

var Public *Mod = new(Mod)
var ActiveMod *Mod

type Mod struct {
	zip *zip.ReadCloser	
}

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
	
	ioutil.WriteFile(Mods+"mod.json", []byte(`{"name": "0ed","version": "0.1","label": "Sample Mod","description": "This is an example discription","dependencies": []} `), 0755)
	
	var err error
	//Try using 0ad's public.zip file.
	// Open a zip archive for reading.
	Public.zip, err = zip.OpenReader("/usr/share/0ad/data/mods/public/public.zip")

		if err != nil {
		
	Public.zip, err = zip.OpenReader(uid.HomeDir+"/AppData/Local/0 A.D. alpha/data/mods/public/public.zip")
	
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
		
		for _, f := range Public.zip.File {
			if f.Name == path {
				return f.Open()
			}
		}
	}
	return nil, errors.New("File not found!")
}

//Only works for public right now.
func (m *Mod) Templates() []string {
	//Retrieve all template files from the 0ad public mod.
	var result []string
	for _, f := range m.zip.File {
		if strings.Contains(f.Name, Templates) {
			name := strings.Split(f.Name, Templates)[1]
			if name[len(name)-4:] == ".xml" {
				result = append(result, name[:len(name)-4])
			}
		}
	}
	sort.Strings(result)
	
	return result
}

//Only works for public right now.
func (m *Mod) Components() []string {
	var result []string
	for _, f := range m.zip.File {
		if strings.Contains(f.Name, Components) {
			name := strings.Split(f.Name, Components)[1]
			if name[len(name)-3:] == ".js" {
				result = append(result, name[:len(name)-3])
			}
		}
	}
	sort.Strings(result)
	
	return result
}
