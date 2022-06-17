package src

var PACKAGE_MANAGERS = []string{
	"apt",
	"pacman",
	"dnf",
	"zypper",
	"apk",
	"rpm",
}

var FLAGS = []string{
	"zip", //bootstrap from which zip file
	"a",   //if true it looks for config.json, the more extensive config for all Operatin Systems
	"c",   //should the config files from the zip be stored in $HOME/.config/conffusion
}

var ALLCONFIG string = "config.json"
var LINCONFIG string = "lconfig.txt"
