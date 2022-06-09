package src

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

func CreateSyncFolder(verbose bool) string {
	DirPath := path.Join(os.TempDir(), "conffusion"+time.Now().String())
	err := os.MkdirAll(DirPath, 0700)
	if err != nil {
		log.Fatalln("Could not create temp folder!")
	}
	if verbose {
		fmt.Printf("Created folder %s\n", DirPath)
	}
	return DirPath

}

func WriteFile(verbose bool, filepath string, content string) string {
	err := ioutil.WriteFile(filepath, []byte(content), 0600)
	if err != nil {
		log.Printf("Could not write file %s\n", filepath)
		return "error"
	}
	if verbose {
		fmt.Printf("Successfully writen %s\n", filepath)
	}
	return "success"
}
