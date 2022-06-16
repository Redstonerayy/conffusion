package src

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//create folder
func CreateFolder(verbose bool, path string) error {
	err := os.MkdirAll(path, 0700)
	if err != nil {
		log.Println("Couldn't create folder!")
		return errors.New("couldn't create folder")
	}
	if verbose {
		fmt.Printf("Created folder %s\n", path)
	}
	return nil

}

//write file
func WriteFile(verbose bool, filepath string, content []byte) error {
	err := ioutil.WriteFile(filepath, content, 0600)
	if err != nil {
		log.Printf("Could not write file %s\n", filepath)
		return errors.New("could not write file " + filepath)
	}
	if verbose {
		fmt.Printf("Successfully writen %s\n", filepath)
	}
	return nil
}

//readfile and return binary data
func ReadFile(verbose bool, filepath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf("Could not read file %s\n", filepath)
		return []byte{}, errors.New("could not read file " + filepath)
	}
	if verbose {
		fmt.Printf("Successfully read file %s\n", filepath)
	}
	return data, nil
}

//read file and write to new file, return error of failing function
func CopyFile(verbose bool, filesource string, filetarget string) error {
	filedata, err := ReadFile(verbose, filesource)
	//write when no error
	if err == nil {
		//write data, converting it to a string before that
		writeerr := WriteFile(verbose, filetarget, filedata)
		return writeerr
	}
	return err
}
