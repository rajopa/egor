package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DB_URL string `mapstructure:"DB_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")    // Ищем в корне проекта
	viper.SetConfigName(".env") // Имя файла
	viper.SetConfigType("env")  // Формат файла

	viper.AutomaticEnv() // Читать системные переменные, если они есть

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}
