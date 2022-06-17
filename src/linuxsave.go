//function to execute on linux
package src

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

func LinuxSave(verbose bool, configfolder string, zipfiles bool, deltefolder bool, configtype string) {
	//read config files
	Variables, vardata := ReadVariables(configfolder, verbose)

	var configdata []byte
	var LinuxConfig []string
	var AllConfig map[string]interface{}

	if configtype == LINCONFIG {
		LinuxConfig, configdata, _ = ReadTxtConfig(configfolder, verbose)
	} else {
		AllConfig, configdata = ReadJsonConfig(configfolder, verbose)
	}

	//create sync folder to write files to and zip in the end
	Homedir, _ := os.UserHomeDir()
	SyncPath := path.Join(Homedir, "conffusionbackup", "conffusion"+time.Now().Format("2006-01-02_3:4:5_pm"))
	syncfoldererr := CreateFolder(verbose, SyncPath)
	if syncfoldererr != nil {
		log.Fatalf("Couldn't create sync folder %s", SyncPath)
	}

	//write config files
	configwriterr := WriteFile(verbose, path.Join(SyncPath, configtype), configdata)
	if configwriterr != nil {
		log.Print("Couldn't write config.json to sync folder!")
	}
	varswriteerr := WriteFile(verbose, path.Join(SyncPath, "vars.txt"), vardata)
	if varswriteerr != nil {
		log.Print("Couldn't write config.json to sync folder!")
	}

	//save package list
	//get packages
	PkgManager, _ := GetPackageManager(verbose)
	SysPackages, _ := GetPackages(verbose, PkgManager)
	//write list
	packagelisterr := WriteFile(verbose, path.Join(SyncPath, "pkgs.txt"), []byte(strings.Join(SysPackages, "\n")))
	if packagelisterr != nil {
		log.Printf("Couldn't write pkg list %s", path.Join(SyncPath, "pkgs.txt"))
	}

	//save files depending on config type
	//type assertions are needed when dealing with this map
	if configtype == LINCONFIG {
		for i := 0; i < len(LinuxConfig); i++ {
			//path for linux
			linux := LinuxConfig[i]
			//replace variables in paths e.g. $USERNAME with their value
			for key, val := range Variables {
				linux = strings.Replace(linux, "$"+key, val, -1)
			}
			//copy config file
			err := CopyFile(false, linux, path.Join(SyncPath, fmt.Sprint(i)+".txt"))
			if err != nil {
				log.Printf("Couldn't copy file %s\n", linux)
			}
		}
	} else {
		configfiles := AllConfig["files"].([]interface{})
		for i := 0; i < len(configfiles); i++ {
			//id of file
			id := configfiles[i].(map[string]interface{})["id"].(float64)
			//path for linux
			linux := configfiles[i].(map[string]interface{})["linux"].(string)
			//replace variables in paths e.g. $USERNAME with their value
			for key, val := range Variables {
				linux = strings.Replace(linux, "$"+key, val, -1)
			}
			//copy config file
			err := CopyFile(false, linux, path.Join(SyncPath, fmt.Sprint(id)+".txt"))
			if err != nil {
				log.Printf("Couldn't copy file %s\n", linux)
			}
		}
	}

	//save config groups
	//TODO

	//add things to do before zipping, e.g. copying other files or executing other scripts
	//TODO

	//zip files in folder eventually
	if zipfiles {
		FilesToZip := []string{}
		FileSystem := os.DirFS(SyncPath)

		fs.WalkDir(FileSystem, ".", func(filepath string, d fs.DirEntry, err error) error {
			if err != nil {
				log.Printf("Error reading %s\n", filepath)
			}
			if filepath != "." {
				FilesToZip = append(FilesToZip, path.Join(SyncPath, filepath))
			}
			return nil
		})

		//zip files
		Path := path.Join(SyncPath + ".zip")
		if err := ZipFiles(Path, FilesToZip); err != nil {
			log.Println("Couldn't zip files!")
		}
		fmt.Println("Zipped File:", Path)
	}

	//delete temporary folder so only zipped one remains, can be changed
	if deltefolder {
		syncdeleteerr := os.RemoveAll(SyncPath)
		if syncdeleteerr != nil {
			log.Printf("Couldn't remove folder %s\n", SyncPath)
		}
	}
}
