package config

type Insurance struct {
	Name           string  `mapstructure:"name" json:"name" yaml:"name"`
	CreditCode     string  `mapstructure:"credit-code" json:"credit-code" yaml:"credit-code"`
	Address        string  `mapstructure:"address" json:"address" yaml:"address"`
	ZipCode        string  `mapstructure:"zip-code" json:"zip-code" yaml:"zip-code"`
	Tel            string  `mapstructure:"tel" json:"tel" yaml:"tel"`
	BankName       string  `mapstructure:"bankName" json:"bankName" yaml:"bankName"`
	BankNo         string  `mapstructure:"bankNo" json:"bankNo" yaml:"bankNo"`
	InsuranceToken string  `mapstructure:"insurance-token" json:"insurance-token" yaml:"insurance-token"`
	HostDomain     string  `mapstructure:"host-domain" json:"host-domain" yaml:"host-domain"`
	ElogPrefix     string  `mapstructure:"elog-prefix" json:"elog-prefix" yaml:"elog-prefix"`
	TempDir        string  `mapstructure:"temp-dir" json:"temp-dir" yaml:"temp-dir"`
	LogoFile       string  `mapstructure:"logo-file" json:"logo-file" yaml:"logo-file"`
	KeyFile        string  `mapstructure:"key-file" json:"key-file" yaml:"key-file"`
	StampFile      string  `mapstructure:"stamp-file" json:"stamp-file" yaml:"stamp-file"`
	LegalFile      string  `mapstructure:"legal-file" json:"legal-file" yaml:"legal-file"`
	SignProgram    string  `mapstructure:"sign-program" json:"sign-program" yaml:"sign-program"`
	APIDomain      string  `mapstructure:"api-domain" json:"api-domain" yaml:"api-domain"`
	JRAPIDomain    string  `mapstructure:"jr-api-domain" json:"jr-api-domain" yaml:"jr-api-domain"`
	AppKey         string  `mapstructure:"app-key" json:"app-key" yaml:"app-key"`
	AppSecret      string  `mapstructure:"app-secret" json:"app-secret" yaml:"app-secret"`
	DiskPath       string  `mapstructure:"disk-path" json:"disk-path" yaml:"disk-path"`
	ElogRate       float64 `mapstructure:"elog-rate" json:"elog-rate" yaml:"elog-rate"`
	ElogMinAmount  float64 `mapstructure:"elog-min-amount" json:"elog-min-amount" yaml:"elog-min-amount"`
	NNRequestUrl   string  `mapstructure:"nn-request-url" json:"nn-request-url" yaml:"nn-request-url"`
	NNAppKey       string  `mapstructure:"nn-app-key" json:"nn-app-key" yaml:"nn-app-key"`
	NNAppSecret    string  `mapstructure:"nn-app-secret" json:"nn-app-secret" yaml:"nn-app-secret"`
	NNTaxNo        string  `mapstructure:"nn-tax-no" json:"nn-tax-no" yaml:"nn-tax-no"`
	NNTaxRate      float64 `mapstructure:"nn-tax-rate" json:"nn-tax-rate" yaml:"nn-tax-rate"`
	NNChecker      string  `mapstructure:"nn-checker" json:"nn-checker" yaml:"nn-checker"`
	NNPayee        string  `mapstructure:"nn-payee" json:"nn-payee" yaml:"nn-payee"`
	NNClerk        string  `mapstructure:"nn-clerk" json:"nn-clerk" yaml:"nn-clerk"`
	NNAccessToken  string  `mapstructure:"nn-access-token" json:"nn-access-token" yaml:"nn-access-token"`
	NNGoodsCode    string  `mapstructure:"nn-goods-code" json:"nn-goods-code" yaml:"nn-goods-code"`
}
