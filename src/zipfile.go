package src

import (
	"archive/zip"
	"errors"
	"io"
	"log"
	"os"
)

// copied from https://golangcode.com/create-zip-files-in-go/
func ZipFiles(filepath string, files []string) error {
	//create file and stack close call
	newZipFile, err := os.Create(filepath)
	if err != nil {
		return errors.New("couln't create zip folder")
	}
	defer newZipFile.Close()

	//create zipwriter and stack close call
	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		if err = AddFileToZip(zipWriter, file); err != nil {
			log.Printf("Couldn't add file %s", file)
		}
	}
	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filepath string) error {
	//open file and stack close call
	fileToZip, err := os.Open(filepath)
	if err != nil {
		return errors.New("could not open file " + filepath)
	}
	defer fileToZip.Close()

	//get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return errors.New("could not get file stats " + filepath)
	}

	//extract header info
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	//change to deflate to gain better compression
	//see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	//copy file into zip
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
