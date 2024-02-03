package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

func Yaml[T any](path string, configuration *T) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(configuration)
	if err != nil {
		return err
	}

	err = envconfig.Process("", configuration)
	if err != nil {
		return err
	}

	return nil
}
