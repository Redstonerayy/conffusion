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
		return "write_error"
	}
	if verbose {
		fmt.Printf("Successfully writen %s\n", filepath)
	}
	return "write_success"
}

func ReadFile(verbose bool, filepath string) ([]byte, string) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf("Could not read file %s\n", filepath)
		return nil, "read_error"
	}
	if verbose {
		fmt.Printf("Successfully read file %s\n", filepath)
	}
	return data, "read_success"
}

func CopyFile(verbose bool, filesource string, filetarget string) string {
	filedata, err := ReadFile(verbose, filesource)
	if err != "read_error" {
		WriteReturn := WriteFile(verbose, filetarget, string(filedata))
		return WriteReturn
	}
	return err
}
