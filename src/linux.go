//function to execute on linux
package src

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

func Linux(verbose bool, configfolder string, zipfiles bool) {
	//read variables
	filedata, err := ReadFile(verbose, path.Join(configfolder, "vars.txt"))
	var Variables = make(map[string]string)
	if err != nil {
		log.Fatalf("Couldn't read configfolder %s", configfolder)
	} else {
		//create map with variables
		for _, i := range strings.Split(string(filedata), "\n") {
			parts := strings.Split(i, " ")
			Variables[parts[0]] = parts[1]
		}
	}

	//read json config file into an unstructed map
	filedata, err = ReadFile(verbose, path.Join(configfolder, "config.json"))
	var result map[string]interface{}
	if err != nil {
		log.Fatalf("Couldn't read configfolder %s", configfolder)
	} else {
		json.Unmarshal([]byte(filedata), &result)
	}

	//create sync folder to write files to and zip in the end
	Homedir, _ := os.UserHomeDir()
	SyncPath := path.Join(Homedir, "conffusionbackup", "conffusion"+time.Now().Format("2006-01-02_3:4:5_pm"))
	syncfoldererr := CreateFolder(verbose, SyncPath)
	if syncfoldererr != nil {
		log.Fatalf("Couldn't create sync folder %s", SyncPath)
	}

	//save package list
	//get packages
	PkgManager, _ := GetPackageManager(verbose)
	SysPackages, _ := GetPackages(verbose, PkgManager)
	//write list
	packagelisterr := WriteFile(verbose, path.Join(SyncPath, "pkgs.txt"), strings.Join(SysPackages, "\n"))
	if packagelisterr != nil {
		log.Printf("Couldn't write pkg list %s", path.Join(SyncPath, "pkgs.txt"))
	}

	//save config files
	//type assertions are needed when dealing with this map
	configfiles := result["files"].([]interface{})
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

	//save config groups
	//TODO

	//add things to do before zipping, e.g. copying other files or executing other scripts
	//TODO

	//zip files in folder eventually zip later
	if zipfiles {
		files, err := os.ReadDir(SyncPath)
		if err != nil {
			log.Printf("Couldn't read folder %s to compress it!\n", SyncPath)
		} else {
			// generate string array of files to zip
			FilesToZip := []string{}
			for _, value := range files {
				FilesToZip = append(FilesToZip, path.Join(SyncPath, value.Name()))
			}

			//zip files
			Path := path.Join(SyncPath + ".zip")
			if err := ZipFiles(Path, FilesToZip); err != nil {
				log.Println("Couldn't zip files!")
			}
			fmt.Println("Zipped File:", Path)
		}
	}
}
