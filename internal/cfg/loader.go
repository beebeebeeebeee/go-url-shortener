package cfg

import (
	"github.com/spf13/viper"
	"log"
)

func LoadEnv(path string) {
	v := viper.NewWithOptions(viper.KeyDelimiter("_"))
	v.AddConfigPath(path)
	v.SetConfigType("dotenv")
	v.SetConfigName(".env")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("error on reading config file, %s", err)
	}

	err := v.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("unable to decode into config, %s", err)
	}

}
