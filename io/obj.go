package io

import (
	"encoding/json"
	"io/ioutil"
)

func ReadFileObj[T any](path string) ([]T, error) {
	_, err := IsExistFile(path)
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

	err = ioutil.WriteFile(path, content, 0644)
	if err != nil {
		return nil
	}
	return nil
}