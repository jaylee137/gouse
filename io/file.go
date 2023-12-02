package io

import (
	"os"
	"bufio"
)

func IsExistFile(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_, err := os.Create(path)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func CreateFile(path string) error {
	_, err := os.Create(path)
	if err != nil {
		return err
	}
	return nil
}

func RemoveFile(path string) error {
	_, err := IsExistFile(path)
	if err != nil {
		return err
	}

	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

func ReadFileByLine(path string) ([]string, error) {
	_, err := IsExistFile(path)
	if err != nil {
		return nil, err
	}


	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func FileInfo(path string) (os.FileInfo, error) {
	_, err := IsExistFile(path)
	if err != nil {
		return nil, err
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return fileInfo, nil
}

func RenameFile(oldPath, newPath string) error {
	_, err := IsExistFile(oldPath)
	if err != nil {
		return err
	}

	err = os.Rename(oldPath, newPath)
	if err != nil {
		return err
	}

	return nil
}