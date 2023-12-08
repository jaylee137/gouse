package io

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func CreatePath(relativePath string) error {
	// Get the absolute path from the relative path
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		return err
	}

	// Extract the directory from the absolute path
	dir := filepath.Dir(absolutePath)

	// Create the directory and any necessary parent directories
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// Create the file
	file, err := os.Create(absolutePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func ReadPath(relativePath string) ([]byte, error) {
	// Get the absolute path from the relative path
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		return nil, err
	}

	// Read the file
	content, err := ioutil.ReadFile(absolutePath)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func WritePath(relativePath string, content []byte) error {
	// Get the absolute path from the relative path
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		return err
	}

	// Write to the file
	err = ioutil.WriteFile(absolutePath, content, 0644)
	if err != nil {
		return err
	}

	return nil
}
