package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"

	Constant "github.com/rizqyfahmi/gin-greetings-clean-architecture/constant"
	CustomErrorPackage "github.com/rizqyfahmi/gin-greetings-clean-architecture/pkg/custom_error"
)

type Environment struct {
	App AppEnvironment
}

type AppEnvironment struct {
	Port           string `env-required:"true" env:"PORT"`
	RequestTimeout int    `env-required:"true" env:"REQUEST_TIMEOUT"`
}

type Config interface {
	Setup() error
	GetConfig() *Environment
}

type ConfigImpl struct {
	environment *Environment
}

func NewConfig() Config {
	return &ConfigImpl{}
}

func (c *ConfigImpl) Setup() error {
	path := "Config:Setup"
	directory, err := os.Getwd()
	if err != nil {
		return CustomErrorPackage.NewCustomError(
			Constant.ErrConfigPath,
			err,
			path,
		)
	}

	var environment Environment
	envFile := fmt.Sprintf("%s/.env", directory)
	err = cleanenv.ReadConfig(envFile, &environment)
	if err != nil {
		return CustomErrorPackage.NewCustomError(
			Constant.ErrConfig,
			err,
			path,
		)
	}

	c.environment = &environment
	return nil
}

func (c *ConfigImpl) GetConfig() *Environment {
	return c.environment
}
