//functions to get info about the operating system or query things
package src

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

//determine operating to check which functions to run
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

//return linux package manager to determine which commands to run
func GetPackageManager(verbose bool) (string, error) {
	ExecFolders := strings.Split(os.Getenv("PATH"), ":")
	//loop over folders in path which could contain the package manager executable
	//package managers are defined in constants.go
	for _, i := range PACKAGE_MANAGERS {
		for _, j := range ExecFolders {
			_, err := os.Stat(path.Join(j, i))
			if err == nil {
				if verbose {
					fmt.Printf("Package Manager found: %s\n", i)
				}
				return i, nil
			}
		}
	}
	log.Println("Package Manager not found!")
	return "", errors.New("package Manager not found")
}

//query a list of packages for a manager to dump to file
//and enable reinstalling them later
func GetPackages(verbose bool, manager string) ([]string, error) {
	//execute listing command
	var CMD *exec.Cmd
	switch manager {
	case "apt":
		CMD = exec.Command("apt", "list", "--installed")
	case "pacman":
		CMD = exec.Command("pacman", "-Q")
	default:
		log.Println("Package Manager not supported!")
		return []string{}, errors.New("package Manager not supported")
	}
	out, err := CMD.Output()
	if err != nil {
		log.Fatalln("Could not query packages!")
		return []string{}, errors.New("Could not query packages!")
	}
	//format output. needs testing for each manager
	OutString := string(out)
	OutLines := strings.Split(OutString, "\n")
	if verbose {
		for i := 0; i < len(OutLines)-1; i++ {
			fmt.Printf("Packages %d: %s \n", i+1, OutLines[i])
		}
	}
	return OutLines, nil
}
