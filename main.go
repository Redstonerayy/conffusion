package main

import src "github.com/Redstonerayy/conffusion/src"

func main() {
	SysOs := src.GetOS(true)
	ConfigFolder := "/home/anton/conffusion/src/groups"
	switch SysOs {
	case "linux":
		src.Linux(true, ConfigFolder)
		// case "windows":
		// 	Windows()
		// case "darwin":
		// 	Darwin()
	}
}
