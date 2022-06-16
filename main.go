/*
   <Backup config files and reuse them or bootstrap new pcs>
   Copyright (C) <2022>  <Anton R.>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"log"
	"os"
	"path"
	"strings"

	src "github.com/Redstonerayy/conffusion/src"
)

func main() {
	verbose := false
	SysOs := src.GetOS(verbose)

	//read main config file at homedir/.conffusion
	HomeDir, _ := os.UserHomeDir()
	EtcData, etcerr := src.ReadFile(verbose, path.Join(HomeDir, ".conffusion"))
	var EtcVars = make(map[string]string)
	if etcerr != nil {
		log.Fatalln("Couldn't read config file!")
	} else {
		//create map with config variables
		for _, i := range strings.Split(string(EtcData), "\n") {
			parts := strings.Split(i, " ")
			EtcVars[parts[0]] = parts[1]
		}
	}

	//parse processs arguments
	Operation, Flags := src.ParseArguments(verbose)
	configtype := src.LINCONFIG
	if _, ok := Flags["a"]; ok {
		configtype = src.ALLCONFIG
	}

	if Operation == "backup" {
		//make backup
		ConfigFolder := EtcVars["CONFIGFOLDER"]
		//execute specific function for each os
		switch SysOs {
		case "linux":
			src.LinuxSave(verbose, ConfigFolder, true, true, configtype)
			// case "windows":
			// 	Windows()
			// case "darwin":
			// 	Darwin()
		}
	} else if Operation == "bootstrap" {
		//make bootstrap
		var ZipFile string
		if val, ok := Flags["zip"]; ok {
			ZipFile = val
		}
		//execute specific function for each os
		switch SysOs {
		case "linux":
			src.LinuxBootstrap(verbose, ZipFile, true, true, configtype)
			// case "windows":
			// 	Windows()
			// case "darwin":
			// 	Darwin()
		}
	}

}
