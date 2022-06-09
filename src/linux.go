package src

import (
	"encoding/json"
	"log"
	"path"
	"strings"
)

func Linux(verbose bool, configfolder string) {
	//read vars
	filedata, err := ReadFile(path.Join(configfolder, "vars.txt"))
	if err != "read_success" {
		log.Printf("Couldn't read configfolder %s", configfolder)
	} else {
		var Variables map[string]string
		for _, i := range strings.Split(string(filedata), "\n") {
			parts := strings.Split(i, " ")
			Variables[parts[0]] = parts[1]
		}
	}
	//read config
	filedata, err = ReadFile(verbose, path.Join(configfolder, "config.json"))
	var jsonerr;
	if err != "read_success" {
		log.Printf("Couldn't read configfolder %s", configfolder)
	} else {
		var config Config
		jsonerr = json.Unmarshal(filedata, &config)
	}
	jsonerr.config.Config.Files.Id
	//create sync folder
	DirPath := CreateSyncFolder(verbose)
	//save package list
	PkgManager := GetPackageManager(verbose)
	SysPackages := GetPackages(verbose, PkgManager)
	_ = WriteFile(verbose, path.Join(DirPath, "pkgs.txt"), strings.Join(SysPackages, "\n"))
	//save config files

}
