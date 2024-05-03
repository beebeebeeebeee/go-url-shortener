package cfg

var Cfg Config

type AppConfig struct {
	Port    int    `mapstructure:"PORT"`
	BaseURL string `mapstructure:"BASEURL"`
}

type Config struct {
	App AppConfig `mapstructure:"APP"`
}
