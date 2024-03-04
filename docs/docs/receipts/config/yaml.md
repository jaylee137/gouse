# Yaml

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/config"
)
```
## Functions


### SampleConfigYaml

```go
func SampleConfigYaml() {
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
