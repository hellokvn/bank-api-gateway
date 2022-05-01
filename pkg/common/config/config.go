package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port               string `mapstructure:"PORT"`
	BankAccountQSvcUrl string `mapstructure:"BANK_ACCOUNT_QUERY_SVC_URL"`
	BankAccountCSvcUrl string `mapstructure:"BANK_ACCOUNT_COMMAND_SVC_URL"`
	BankFundsQSvcUrl   string `mapstructure:"BANK_FUNDS_QUERY_SVC_URL"`
	BankFundsCSvcUrl   string `mapstructure:"BANK_FUNDS_COMMAND_SVC_URL"`
	IsDOcker           string `mapstructure:"IS_DOCKER"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./")

	path := ".env"

	if os.Getenv("IS_DOCKER") == "true" || os.Getenv("IS_DOCKER") == "1" {
		path = ".docker.env"
	}

	viper.SetConfigName(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
