package config

import (
	"github.com/spf13/cast"
)

const (
	varLogLevel = "log.level"
	varHTTPPort = "http.port"
)

type Config struct {
	config map[string]interface{}
}

func New() *Config {
	c := new(Config)
	c.config = make(map[string]interface{})
	c.setDefaults()
	return c
}

func (c *Config) setDefaults() {
	c.config[varLogLevel] = "debug"
	c.config[varHTTPPort] = 8080
}

func (c *Config) GetLogLevel() string {
	return cast.ToString(c.config[varLogLevel])
}

func (c *Config) GetHttpPort() int {
	return cast.ToInt(c.config[varHTTPPort])
}
