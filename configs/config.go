package config

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
)

type config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUsername    string `mapstructure:"DB_USERNAME"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	WebServerPort string `mapstructure:"WEBSERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadEnv(path string) (config, error) {
	var cfg config
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigName("GoApi")
	viper.AutomaticEnv()
	err := viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, err
}
