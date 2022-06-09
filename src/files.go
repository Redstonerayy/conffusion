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
	err := os.MkdirAll(DirPath, 700)
	if err != nil {
		log.Fatalln("Could not create temp folder!")
	}
	if verbose {
		fmt.Printf("Created folder %s", DirPath)
	}
	return DirPath

}

func WriteFile(verbose bool, filepath string, content string) {
	err := ioutil.WriteFile(filepath, []byte(content), 600)
	if err != nil {
		log.Printf("Could not write file %s\n", filepath)
	}
}
