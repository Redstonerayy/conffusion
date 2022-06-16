package src

import (
	"encoding/json"
	"errors"
	"log"
	"path"
	"strings"
)

func ReadVariables(configfolder string, verbose bool) (map[string]string, []byte) {
	//read variables
	vardata, varerr := ReadFile(verbose, path.Join(configfolder, "vars.txt"))
	var Variables = make(map[string]string)
	if varerr != nil {
		log.Fatalf("Couldn't read configfolder %s", configfolder)
	} else {
		//create map with variables
		for _, i := range strings.Split(string(vardata), "\n") {
			parts := strings.Split(i, " ")
			Variables[parts[0]] = parts[1]
		}
	}
	return Variables, vardata
}

func ReadJsonConfig(configfolder string, verbose bool) (map[string]interface{}, []byte) {
	//read json config file into an unstructed map
	configdata, configerr := ReadFile(verbose, path.Join(configfolder, ALLCONFIG))
	var Config map[string]interface{}
	if configerr != nil {
		log.Fatalf("Couldn't read configfolder %s", configfolder)
	} else {
		json.Unmarshal([]byte(configdata), &Config)
	}
	return Config, configdata
}

func ReadTxtConfig(configfolder string, verbose bool) ([]string, []byte, error) {
	//read json config file into an unstructed map
	configdata, configerr := ReadFile(verbose, path.Join(configfolder, LINCONFIG))
	if configerr != nil {
		log.Fatalf("Couldn't read configfolder %s", configfolder)
		return []string{}, configdata, errors.New("could not read configfolder" + configfolder)
	} else {
		Config := strings.Split(string(configdata), "\n")
		return Config, configdata, nil
	}
}
