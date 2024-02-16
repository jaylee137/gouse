package config

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/config"
)

func configJson() {
	type Configuration struct {
		Users  []string
		Groups []string
	}

	var myConf Configuration
	err := config.Json("conf.json", &myConf)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Users:", myConf.Users)
	fmt.Println("Groups:", myConf.Groups)
}
