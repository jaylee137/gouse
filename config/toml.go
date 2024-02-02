package config

import (
	"github.com/pelletier/go-toml"
)

func Toml[T any](path string, configuration *T) error {
	data, err := toml.LoadFile(path)
	if err != nil {
		return err
	}

	err = data.Unmarshal(configuration)
	if err != nil {
		return err
	}

	return nil
}
