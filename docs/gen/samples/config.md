# Config

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/config"
)
```
## Functions


### configJson

```go
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
```

### configYaml

```go
func configYaml() {
	type Configuration struct {
		Server struct {
			Port string `envconfig:"SERVER_PORT"`
			Host string `envconfig:"SERVER_HOST"`
		}
		Database struct {
			Username string `envconfig:"DATABASE_USERNAME"`
			Password string `envconfig:"DATABASE_PASSWORD"`
		}
	}

	var myConf Configuration
	err := config.Yaml("conf.yaml", &myConf)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Server Port:", myConf.Server.Port)
	fmt.Println("Server Host:", myConf.Server.Host)
	fmt.Println("DB Username:", myConf.Database.Username)
	fmt.Println("DB Password:", myConf.Database.Password)
}
```

### configToml

```go
func configToml() {
	type Configuration struct {
		Mysql struct {
			Host     string
			Username string
			Password string
			Database string
		}
	}

	var myConf Configuration
	err := config.Toml("conf.toml", &myConf)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Mysql Host:", myConf.Mysql.Host)
	fmt.Println("Mysql Username:", myConf.Mysql.Username)
	fmt.Println("Mysql Password:", myConf.Mysql.Password)
	fmt.Println("Mysql Database:", myConf.Mysql.Database)
}
```
