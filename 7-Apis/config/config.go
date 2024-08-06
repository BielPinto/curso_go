package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBName       string `mapstructure:"DB_NAME"`
	DBServerPort string `mapstructure:"DB_SERVER_PORT"`
	JWTSecret    string `mapstructure:"DB_SECRET"`
	DBExpiresIn  string `mapstructure:"DB_EXPIRESIN"`
	TokenAuth    *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {

	viper.SetConfigName("app_config")
	viper.SetConfigType("emv")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, err
}
