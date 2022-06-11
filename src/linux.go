package src

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func Linux(verbose bool, configfolder string, zipfiles bool) {
	//read vars
	filedata, err := ReadFile(verbose, path.Join(configfolder, "vars.txt"))
	var Variables = make(map[string]string)
	if err != "read_success" {
		log.Printf("Couldn't read configfolder %s", configfolder)
	} else {
		for _, i := range strings.Split(string(filedata), "\n") {
			parts := strings.Split(i, " ")
			Variables[parts[0]] = parts[1]
		}
	}
	//read config
	filedata, err = ReadFile(verbose, path.Join(configfolder, "config.json"))
	var result map[string]interface{}
	if err != "read_success" {
		log.Printf("Couldn't read configfolder %s", configfolder)
	} else {
		json.Unmarshal([]byte(filedata), &result)
	}
	//create sync folder
	DirPath := CreateSyncFolder(verbose, true)
	//save package list
	PkgManager := GetPackageManager(verbose)
	SysPackages := GetPackages(verbose, PkgManager)
	_ = WriteFile(verbose, path.Join(DirPath, "pkgs.txt"), strings.Join(SysPackages, "\n"))
	//save config files
	for i := 0; i < len(result["files"].([]interface{})); i++ {
		id := result["files"].([]interface{})[i].(map[string]interface{})["id"].(float64)
		linux := result["files"].([]interface{})[i].(map[string]interface{})["linux"].(string)
		for key, val := range Variables {
			linux = strings.Replace(linux, "$"+key, val, -1)
		}
		err := CopyFile(false, linux, path.Join(DirPath, fmt.Sprint(id)+".txt"))
		if err == "write_error" {
			log.Printf("Couldn't copy file %s\n", linux)
		}
	}
	//save config groups

	//zip files in folder eventually zip later
	if zipfiles {
		files, err := os.ReadDir(DirPath)
		if err != nil {
			log.Printf("Couldn't read folder %s to compress it!\n", DirPath)
		} else {
			// generate string array of files to zip
			FilesToZip := []string{}
			for _, value := range files {
				FilesToZip = append(FilesToZip, path.Join(DirPath, value.Name()))
				fmt.Println(FilesToZip)
			}

			Path := path.Join(DirPath + ".zip")
			if err := ZipFiles(Path, FilesToZip); err != nil {
				log.Println("Couldn't zip files!")
			}
			fmt.Println("Zipped File:", Path)
		}
	}
}
