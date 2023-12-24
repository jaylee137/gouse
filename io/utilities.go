package io

import (
	"io"
	"path/filepath"

	"archive/zip"
	"encoding/base64"
	"os"
)

func Zip(zipFileName string, files []string) error {
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		fileToZip, err := os.Open(file)
		if err != nil {
			return err
		}
		defer fileToZip.Close()

		zipWriter, err := zipWriter.Create(file)
		if err != nil {
			return err
		}

		_, err = io.Copy(zipWriter, fileToZip)
		if err != nil {
			return err
		}
	}
	return nil
}

func Unzip(zipFile, destFolder string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

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

		_, err = io.Copy(targetFile, fileToExtract)
		if err != nil {
			return err
		}
	}

	return nil
}

func Encode(input []byte) ([]byte, error) {
	encodedData := base64.StdEncoding.EncodeToString(input)
	return []byte(encodedData), nil
}

func Decode(input []byte) ([]byte, error) {
	decodedData, err := base64.StdEncoding.DecodeString(string(input))
	if err != nil {
		return nil, err
	}
	return decodedData, nil
}
