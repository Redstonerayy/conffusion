package src

import (
	"encoding/json"
	"fmt"
	"log"
	"path"
	"strings"
)

func Linux(verbose bool, configfolder string) {
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
		fmt.Println(result["files"])
	}
	//create sync folder
	DirPath := CreateSyncFolder(verbose)
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
}
