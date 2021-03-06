package config

import (
	"fmt"
	"path"

	"github.com/goava/di"
	"github.com/urionz/config"
	"github.com/urionz/config/hcl"
	"github.com/urionz/config/ini"
	"github.com/urionz/config/json"
	"github.com/urionz/config/toml"
	"github.com/urionz/config/yaml"
	"github.com/urionz/goofy"
	"github.com/urionz/ini/dotenv"
)

func NewServiceProvider(app goofy.IApplication) {
	app.Provide(func() (*Configure, error) {
		serve = &Configure{
			Config: config.New("goofy"),
		}

		serve.AddDriver(yaml.Driver)
		serve.AddDriver(json.Driver)
		serve.AddDriver(ini.Driver)
		serve.AddDriver(hcl.Driver)
		serve.AddDriver(toml.Driver)

		dotenv.LoadExists(app.Workspace(), ".env")

		envConfFile := dotenv.Get("APP_CONF", fmt.Sprintf("config.%s.toml", dotenv.Get("APP_ENV", "dev")))

		if err := serve.LoadExists(path.Join(app.Workspace(), "config.toml"), path.Join(app.Workspace(), envConfFile)); err != nil {
			return nil, err
		}

		return serve, nil
	}, di.As(new(IConfig)))
}
