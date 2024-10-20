package config

import "github.com/spf13/viper"

type Config struct {
	Mode  string      `mapstructure:"SERVER_MODE"`
	Port  string      `mapstructure:"SERVER_PORT"`
	Mysql ConfigMysql `mapstructure:",squash"`
}

type ConfigMysql struct {
	Host         string `mapstructure:"MYSQL_HOST"`
	Name         string `mapstructure:"MYSQL_DATABASE"`
	Username     string `mapstructure:"MYSQL_USERNAME"`
	Password     string `mapstructure:"MYSQL_PASSWORD"`
	MigrateTable bool   `mapstructure:"MYSQL_MIGRATE_TABLE"`
	ImportData   bool   `mapstructure:"MYSQL_IMPORT_DATA"`
}

func Load(path, name, xten string) (*Config, error) {
	resp := Config{}
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(xten)
	if xten == "env" {
		viper.AutomaticEnv()
	}
	if erro := viper.ReadInConfig(); erro != nil {
		return nil, erro
	}
	if erro := viper.Unmarshal(&resp); erro != nil {
		return nil, erro
	}
	return &resp, nil
}
