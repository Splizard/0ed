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
import "path/filepath"

var Public *Mod = new(Mod)
var ActiveMod *Mod = new(Mod)

var Pyrogenesis string = "pyrogenesis"

type Mod struct {
	Name string
	zip *zip.ReadCloser	
}

var Mods string = "./mods"

func ClosePublic() {
	Public.zip.Close()
}

func ReloadPublic() {
	uid, _ := user.Current()
	
	var err error
	
	//Try using 0ad's public.zip file.
	// Open a zip archive for reading.
	Public.zip, err = zip.OpenReader("/usr/share/0ad/data/mods/public/public.zip")

		if err != nil {
		
	Public.zip, err = zip.OpenReader(uid.HomeDir+"/AppData/Local/0 A.D. alpha/binaries/data/mods/public/public.zip")
	
		if err != nil {
			println("Could not locate 0ad data file!")
		} else {
			Pyrogenesis = uid.HomeDir+"AppData/Local/0 A.D. alpha/binaries/system/pyrogenisis.exe"
		}
	}
}

func init() {
	ActiveMod.Name = "0ed"	
	
	ReloadPublic()
}

func OtherMods() []string {
	files, err := ioutil.ReadDir(Mods+"../")
    if err != nil {
        //Damn
    }
    
    result := []string{}

    for _, f := range files {
		if f.IsDir() && f.Name() != ActiveMod.Name {
			result = append(result, f.Name())
		}
    }
    
    return result
}

func NewMod(name string) *Mod {
	var m = new(Mod)
	m.Name = name
	m.Load()
	return m
}

func (m *Mod) Load() {
	uid, _ := user.Current()
	
	if runtime.GOOS == "linux" {
		Mods = uid.HomeDir+"/.local/share/0ad/mods/"+m.Name+"/"
	}
	
	if runtime.GOOS == "windows" {
		Mods = uid.HomeDir+"/Documents/My Games/0ad/mods/"+m.Name+"/"
	}
	
	os.MkdirAll(Mods, 0755)
	
	ioutil.WriteFile(Mods+"mod.json", []byte(`{"name": "`+m.Name+`","version": "0.1","label": "Sample Mod","description": "This is an example discription","dependencies": []} `), 0755)
	
	ActiveMod = m
}

func (m *Mod) Open(path string) (io.ReadCloser, error) {
	if file, err := os.Open(Mods+path); err == nil {
		return file, err
	}
	
	for _, f := range Public.zip.File {
		if f.Name == path {
			return f.Open()
		}
	}
	return nil, errors.New(path+" not found!")
}

//Only works for public right now.
func (m *Mod) Templates() []string {
	//Retrieve all template files from the 0ad public mod.
	var result []string
	
	if m.zip != nil {
		
		for _, f := range m.zip.File {
			if strings.Contains(f.Name, Templates) {
				name := strings.Split(f.Name, Templates)[1]
				if name[len(name)-4:] == ".xml" {
					result = append(result, name[:len(name)-4])
				}
			}
		}
		
	} else {
		
		filepath.Walk(Mods+Templates, func(path string, info os.FileInfo, err error) error {
			if strings.Contains(path, Templates) {
				name := strings.Split(path, Templates)[1]
				if len(name) > 4 && name[len(name)-4:] == ".xml" {
					result = append(result, name[:len(name)-4])
				}
			}
			return nil
		})
		
	}
	
	sort.Strings(result)
	
	return result
}

//Only works for public right now.
func (m *Mod) Images() []string {
	//Retrieve all template files from the 0ad public mod.
	var result []string
	for _, f := range m.zip.File {
		name := f.Name
		if name[len(name)-4:] == ".dds" {
			result = append(result, name)
		}
	}
	sort.Strings(result)
	
	return result
}

//Only works for public right now.
func (m *Mod) Components() []string {
	var result []string
	
	if m.zip != nil {
		
		for _, f := range m.zip.File {
			if strings.Contains(f.Name, Components) {
				name := strings.Split(f.Name, Components)[1]
				if name[len(name)-3:] == ".js" {
					result = append(result, name[:len(name)-3])
				}
			}
		}
		
	
	} else {
		
		filepath.Walk(Mods+Components, func(path string, info os.FileInfo, err error) error {
			if strings.Contains(path, Components) {
				name := strings.Split(path, Components)[1]
				if len(name) > 3 && name[len(name)-3:] == ".js" {
					result = append(result, name[:len(name)-3])
				}
			}
			return nil
		})
		
	}
	
	sort.Strings(result)
	return result
}
