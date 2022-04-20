package config

import "github.com/spf13/viper"

type Config struct {
	Port               string `mapstructure:"PORT"`
	BankAccountQSvcUrl string `mapstructure:"BANK_ACCOUNT_QUERY_SVC_URL"`
	BankAccountCSvcUrl string `mapstructure:"BANK_ACCOUNT_COMMAND_SVC_URL"`
	BankFundsQSvcUrl   string `mapstructure:"BANK_FUNDS_QUERY_SVC_URL"`
	BankFundsCSvcUrl   string `mapstructure:"BANK_FUNDS_COMMAND_SVC_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
