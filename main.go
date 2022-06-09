package main

import src "github.com/Redstonerayy/conffusion/src"

func main() {
	SysOs := src.GetOS(true)
	switch SysOs {
	case "linux":
		Linux()
	case "windows":
		Windows()
	case "darwin":
		Darwin()
	}
}
