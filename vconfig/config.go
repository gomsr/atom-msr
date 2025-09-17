package vconfig

import (
	"github.com/gomsr/atom-gorm/gconfig"
	"github.com/gomsr/atom-kits/emailx"
	"github.com/gomsr/atom-rest/aconfig"
	"github.com/gomsr/atom-s3/configs3"
	"github.com/gomsr/atom-zap/configz"
	"github.com/kongmsr/oneid-core/modelo"
)

type Server struct {
	System System       `mapstructure:"system" json:"system" yaml:"system"`
	Email  emailx.Email `mapstructure:"email" json:"email" yaml:"email"`
	Zap    configz.Zap  `mapstructure:"zap" json:"zap" yaml:"zap"`

	gconfig.DbServer `yaml:",inline" mapstructure:",squash"`

	// 跨域配置
	Cors      aconfig.CORS     `mapstructure:"cors" json:"cors" yaml:"cors"`
	JWT       aconfig.JWT      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Captcha   aconfig.Captcha  `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	OneidConf modelo.OneidConf `mapstructure:"oneid-token" json:"oneid-token" yaml:"oneid-token"`

	// oss
	Local configs3.Local `mapstructure:"local" json:"local" yaml:"local"`
	AwsS3 configs3.AwsS3 `mapstructure:"aws-s3" json:"aws-s3" yaml:"aws-s3"` // minio|cloudflare same with this
}

type System struct {
	Env          string `mapstructure:"env" json:"env" yaml:"env"`                      // 环境值
	JasyptPwd    string `mapstructure:"jasypt-pwd" json:"jasypt-pwd" yaml:"jasypt-pwd"` // 配置加解密值
	Addr         int    `mapstructure:"addr" json:"addr" yaml:"addr"`                   // 端口值
	UseRedis     bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`    // 使用redis
	LimitType    int    `mapstructure:"limit-type" json:"limit-type" yaml:"limit-type"` // 限流类型:0(本地)|1(本地)|2(redis)
	LimitCountIP int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP  int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	OssType      string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"` // Oss类型
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	IPQueryKey   string `mapstructure:"ip-query-key" json:"ip-query-key" yaml:"ip-query-key"`
	CurrencyKey  string `mapstructure:"currency-key" json:"currency-key" yaml:"currency-key"`
}
