package config

type Insurance struct {
	Name          string  `mapstructure:"name" json:"name" yaml:"name"`
	CreditCode    string  `mapstructure:"credit-code" json:"credit-code" yaml:"credit-code"`
	Address       string  `mapstructure:"address" json:"address" yaml:"address"`
	ZipCode       string  `mapstructure:"zip-code" json:"zip-code" yaml:"zip-code"`
	Tel           string  `mapstructure:"tel" json:"tel" yaml:"tel"`
	TempDir       string  `mapstructure:"temp-dir" json:"temp-dir" yaml:"temp-dir"`
	KeyFile       string  `mapstructure:"key-file" json:"key-file" yaml:"key-file"`
	StampFile     string  `mapstructure:"stamp-file" json:"stamp-file" yaml:"stamp-file"`
	SignProgram   string  `mapstructure:"sign-program" json:"sign-program" yaml:"sign-program"`
	APIDomain     string  `mapstructure:"api-domain" json:"api-domain" yaml:"api-domain"`
	JRAPIDomain   string  `mapstructure:"jr-api-domain" json:"jr-api-domain" yaml:"jr-api-domain"`
	AppKey        string  `mapstructure:"app-key" json:"app-key" yaml:"app-key"`
	AppSecret     string  `mapstructure:"app-secret" json:"app-secret" yaml:"app-secret"`
	ElogRate      float64 `mapstructure:"elog-rate" json:"elog-rate" yaml:"elog-rate"`
	ElogMinAmount float64 `mapstructure:"elog-min-amount" json:"elog-min-amount" yaml:"elog-min-amount"`
}