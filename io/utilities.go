package io

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Zip(zipFileName string, files []string) error {
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := zipFile.Close(); closeErr != nil {
			fmt.Println("Error closing zip file:", closeErr)
		}
	}()

	zipWriter := zip.NewWriter(zipFile)
	defer func() {
		if closeErr := zipWriter.Close(); closeErr != nil {
			fmt.Println("Error closing zip writer:", closeErr)
		}
	}()

	for _, file := range files {
		fileToZip, err := os.Open(file)
		if err != nil {
			return err
		}

		zipEntryWriter, err := zipWriter.Create(file)
		if err != nil {
			fileToZip.Close() // Close the file in case of an error
			return err
		}

		_, err = io.Copy(zipEntryWriter, fileToZip)
		fileToZip.Close() // Close the file after copying
		if err != nil {
			return err
		}
	}

	return nil
}

const maxFileSize = 100 * 1024 * 1024  // 100 MB
const maxTotalSize = 500 * 1024 * 1024 // 500 MB

func Unzip(zipFile, destFolder string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := zipReader.Close(); closeErr != nil {
			fmt.Println("Error closing zip reader:", closeErr)
		}
	}()

	var totalSize int64

	for _, file := range zipReader.File {
		filePath := filepath.Join(destFolder, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		fileToExtract, err := file.Open()
		if err != nil {
			return err
		}
		defer fileToExtract.Close()

		targetFile, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer targetFile.Close()

		// Check the size of the file before copying
		fileSize, err := io.CopyN(targetFile, fileToExtract, maxFileSize)
		if err != nil && err != io.EOF {
			return err
		}

		// Update the total size
		totalSize += fileSize

		// Check the total size limit
		if totalSize > maxTotalSize {
			return errors.New("exceeded maximum total size limit for extraction")
		}
	}

	return nil
}
