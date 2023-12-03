package io

import (
	"bufio"
	"encoding/json"
	"io"
	"os"

	"github.com/thuongtruong1009/gouse/date"
	"github.com/thuongtruong1009/gouse/strings"
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

func WriteFile(path string, data []string) error {
	os.WriteFile(path, []byte(strings.Join(data, " ")), 0600)
	return nil
}

func AppendFile(path string, data []string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return nil
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	for _, v := range data {
		if _, err := file.WriteString(v + "\n"); err != nil {
			return err
		}
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

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

type FileInfoStruct struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime string
	IsDir   bool
	Sys     interface{}
	All     os.FileInfo
}

func FileInfo(path string) (*FileInfoStruct, error) {
	_, err := IsExistFile(path)
	if err != nil {
		return nil, err
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return &FileInfoStruct{
		Name:    fileInfo.Name(),
		Size:    fileInfo.Size(),
		Mode:    fileInfo.Mode(),
		ModTime: date.ISO(fileInfo.ModTime()),
		IsDir:   fileInfo.IsDir(),
		Sys:     fileInfo.Sys(),
		All:     fileInfo,
	}, nil
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

func CopyFile(oldPath, newPath string) error {
	_, err := IsExistFile(oldPath)
	if err != nil {
		return err
	}

	source, err := os.Open(oldPath)
	if err != nil {
		return err
	}

	defer func() {
		if err := source.Close(); err != nil {
			panic(err)
		}
	}()

	target, targetError := os.OpenFile(newPath, os.O_RDWR|os.O_CREATE, 0666)
	if targetError != nil {
		return targetError
	}

	defer func() {
		if err := target.Close(); err != nil {
			panic(err)
		}
	}()

	_, copyError := io.Copy(target, source)
	if copyError != nil {
		return copyError
	}

	return nil
}

func TruncateFile(path string, size int64) error {
	_, err := IsExistFile(path)
	if err != nil {
		return err
	}

	err = os.Truncate(path, size)
	if err != nil {
		return err
	}

	return nil
}

func CleanFile(path string) error {
	TruncateFile(path, 0)
	return nil
}

func ReadFileObj[T any](path string) ([]T, error) {
	_, err := IsExistFile(path)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		return nil, nil
	}

	var data []T

	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func WriteFileObj[T any](path string, data T) error {
	_, err := IsExistFile(path)
	if err != nil {
		err2 := CreateDir(path)
		if err2 != nil {
			return err2
		}
	}

	content, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	err = os.WriteFile(path, content, 0644)
	if err != nil {
		return nil
	}
	return nil
}
