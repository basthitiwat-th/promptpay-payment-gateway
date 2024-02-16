package configs

import "time"

type Config struct {
	MySQL           MySQL
	PromptPayClient PromptPayClient
	MerchantClient  MerchantClient
	Secrets         Secrets
}

type PromptPayClient struct {
	BaseURL     string        `mapstructure:"baseurl" validate:"required"`
	MaxConns    int           `mapstructure:"maxconns" validate:"gt=0"`
	Timeout     time.Duration `mapstructure:"timeout" validate:"gt=0"`
	URLValidate string        `mapstructure:"urlvalidate"`
	URLConfirm  string        `mapstructure:"urlconfirm"`
}

type MerchantClient struct {
	BaseURL    string        `mapstructure:"baseurl" validate:"required"`
	MaxConns   int           `mapstructure:"maxconns" validate:"gt=0"`
	Timeout    time.Duration `mapstructure:"timeout" validate:"gt=0"`
	URLConfirm string        `mapstructure:"urlconfirm"`
}

type MySQL struct {
	Host         string        `mapstructure:"host" validate:"required"`
	Database     string        `mapstructure:"database" validate:"required"`
	Timeout      time.Duration `mapstructure:"timeout" validate:"gt=0"`
	MaxOpenConns *int          `mapstructure:"maxconns"`
	MaxLifetime  time.Duration `mapstructure:"maxlifetime"`
	MaxIdleConns time.Duration `mapstructure:"maxidleconns"`
}

type Secrets struct {
	MySqlUsername string `mapstructure:"mysqlUsername" validate:"required"`
	MySqlPassword string `mapstructure:"mysqlPassword" validate:"required"`
}
