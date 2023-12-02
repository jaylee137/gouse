package io

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func IsExistFile(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_, err := os.Create(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateFile(path string) error {
	_, err := os.Create(path)
	if err != nil {
		return err
	}
	return nil
}

func RemoveFile(path string) error {
	err := IsExistFile(path)
	if err != nil {
		return err
	}

	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile[T any](path string) ([]T, error) {
	err := IsExistFile(path)
    if err != nil {
        return nil, err
    }

	content, err := ioutil.ReadFile(path)
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

func WriteFile[T any](path string, data T) error {
	err := IsExistFile(path)
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

	err = ioutil.WriteFile(path, content, 0644)
	if err != nil {
		return nil
	}
	return nil
}

func CreateDir(path string) error {
	err := IsExistFile(path)
	if err != nil {
		return err
	}

	err = os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}

	return nil
}

func DeleteDir(path string) error {
	err := IsExistFile(path)
	if err != nil {
		return err
	}

	err = os.RemoveAll(path)
	if err != nil {
		return err
	}

	return nil
}
