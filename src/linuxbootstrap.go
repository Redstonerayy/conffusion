package src

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

func LinuxBootstrap(verbose bool, zipfile string, zipfiles bool, deltefolder bool, configtype string, writeconfig bool) {
	// TmpPath := path.Join(os.TempDir(), "temp"+time.Now().Format("2006-01-02_3:4:5_pm"))
	files, _ := ReadZipFile(zipfile)

	var Variables = make(map[string]string)
	//create map with variables
	for _, i := range strings.Split(files["vars.txt"], "\n") {
		parts := strings.Split(i, " ")
		Variables[parts[0]] = parts[1]
	}

	//install package list
	//TODO	PkgManager, _ := GetPackageManager(verbose)
	//write config file
	if writeconfig {
		HomeDir, _ := os.UserHomeDir()
		DefaultConfigPath := path.Join(HomeDir, ".config", "conffusion")
		//create folder if not present
		CreateFolder(verbose, DefaultConfigPath)
		//write config
		_, conferr := os.Stat(path.Join(DefaultConfigPath, configtype))
		if conferr != nil { //file not present
			WriteFile(verbose, path.Join(DefaultConfigPath, configtype), []byte(files[configtype]))
		}
		//write vars
		_, varerr := os.Stat(path.Join(DefaultConfigPath, "vars.txt"))
		if varerr != nil { //file not present
			WriteFile(verbose, path.Join(DefaultConfigPath, "vars.txt"), []byte(files["vars.txt"]))
		}
	}

	//read config and write files depending on config file
	if configtype == LINCONFIG {
		ConfigTxt := strings.Split(files[LINCONFIG], "\n")
		for i := 0; i < len(ConfigTxt); i++ {
			//path for linux
			linux := ConfigTxt[i]
			//replace variables in paths e.g. $USERNAME with their value
			for key, val := range Variables {
				linux = strings.Replace(linux, "$"+key, val, -1)
			}
			//write files
			err := WriteFile(verbose, linux, []byte(files[fmt.Sprint(i)+".txt"]))
			if err != nil {
				log.Printf("Couldn't write file %s\n", linux)
			}
		}
	} else {
		//read config files
		var Config map[string]interface{}
		json.Unmarshal([]byte(files["config.json"]), &Config)

		//type assertions are needed when dealing with this map
		configfiles := Config["files"].([]interface{})
		for i := 0; i < len(configfiles); i++ {
			//id of file
			id := configfiles[i].(map[string]interface{})["id"].(float64)
			//path for linux
			linux := configfiles[i].(map[string]interface{})["linux"].(string)
			//replace variables in paths e.g. $USERNAME with their value
			for key, val := range Variables {
				linux = strings.Replace(linux, "$"+key, val, -1)
			}
			//write files
			err := WriteFile(verbose, linux, []byte(files[strconv.FormatFloat(id, 'f', 0, 64)+".txt"]))
			if err != nil {
				log.Printf("Couldn't write file %s\n", linux)
			}
		}
	}
}
