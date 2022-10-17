package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Config interface {
	Get(key string) string
}

type ConfigImpl struct {
}

func (config *ConfigImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	if err != nil {
		panic(err)
	}

	return &ConfigImpl{}
}
