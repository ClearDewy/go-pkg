/**
 * @ Author: ClearDewy
 * @ Desc:
 **/
package config

type Config struct {
}

func (c *Config) LoadEnvDefault() error {
	return loadEnvDefault(c)
}
