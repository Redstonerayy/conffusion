package src

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

func GetOS(verbose bool) string {
	SystemOs := runtime.GOOS
	if verbose {
		switch SystemOs {
		case "darwin":
			fmt.Println("MacOs")
		case "windows":
			fmt.Println("Windows")
		case "linux":
			fmt.Println("Linux")
		default:
			fmt.Println(SystemOs)
		}
	}
	return SystemOs
}

func GetPackageManager(verbose bool) string {
	ExecFolders := strings.Split(os.Getenv("PATH"), ":")
	for _, i := range PACKAGE_MANAGERS {
		for _, j := range ExecFolders {
			_, err := os.Stat(path.Join(j, i))
			if err == nil {
				if verbose {
					fmt.Printf("Package Manager found: %s\n", i)
				}
				return i
			}
		}
	}
	if verbose {
		log.Fatalln("Package Manager not found!")
	}
	return ""
}

func GetPackages(verbose bool, manager string) []string {
	var CMD *exec.Cmd
	switch manager {
	case "apt":
		CMD = exec.Command("apt", "list", "--installed")
	case "pacman":
		CMD = exec.Command("pacman", "-Q")
	default:
		log.Fatalln("No Package Manager found!")
		return nil
	}
	out, err := CMD.Output()
	if err != nil {
		log.Fatalln("Could not query packages, fatal error, exit!")
	}
	OutString := string(out)
	OutLines := strings.Split(OutString, "\n")
	if verbose {
		for i := 0; i < len(OutLines)-1; i++ {
			fmt.Printf("Packages %d: %s \n", i+1, OutLines[i])
		}
	}
	return OutLines
}
