package config

import (
	"encoding/json"
	"os"
)

func Json[T any](path string, configuration *T) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	decoder := json.NewDecoder(file)
	return decoder.Decode(configuration)
}
