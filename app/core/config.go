package core

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// ConfigMage 全局配置 与开发/生产环境无关
type ConfigMage struct {
	IsProduction bool `yaml:"IsProduction"`
}

type ConfigSite struct {
	Name          string `yaml:"Name"`
	Host          string `yaml:"Host"`
	Https         bool   `yaml:"Https"`
	DeveloperMode bool   `yaml:"DeveloperMode"`
}
type ConfigMysqlConfig struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Dbname   string `yaml:"Dbname"`
}
type ConfigMysql struct {
	Read  ConfigMysqlConfig `yaml:"Read"`
	Write ConfigMysqlConfig `yaml:"Write"`
}

type ConfigSession struct {
	Storage  string `yaml:"Storage"`
	Key      string `yaml:"Key"`
	Lifetime int    `yaml:"Lifetime"`
}
type ConfigSessionRedis struct {
	Host     string  `yaml:"Host"`
	Port     int     `yaml:"Port"`
	Password string  `yaml:"Password"`
	Timeout  float32 `yaml:"Timeout"`
	Maxsize  int     `yaml:"Maxsize"`
	Db       int     `yaml:"Db"`
}

type ConfigSearch struct {
	Storage string `yaml:"Storage"`
}

type ConfigSearchSolr struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
}

type ConfigSearchElasticsearch struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
}

type ConfigMedia struct {
	Storage string `yaml:"Storage"`
}

type ConfigMediaLocal struct {
	Path string `yaml:"Path"`
}

// Config 配置管理和配置文件解析
type Config struct {
	Mage                ConfigMage
	Site                ConfigSite                `yaml:"Site"`
	Mysql               ConfigMysql               `yaml:"Mysql"`
	Session             ConfigSession             `yaml:"Session"`
	SessionRedis        ConfigSessionRedis        `yaml:"SessionRedis"`
	Search              ConfigSearch              `yaml:"Search"`
	SearchElasticsearch ConfigSearchElasticsearch `yaml:"SearchElasticsearch"`
	Media               ConfigMedia               `yaml:"Media"`
	MediaLocal          ConfigMediaLocal          `yaml:"MediaLocal"`
}

// LoadBase 加载全局基础配置
func (config *Config) LoadBase(options map[string]interface{}) {
	mageConfig := &ConfigMage{}
	pathMage := "/config/mage.yaml"
	c0, err := ioutil.ReadFile(pathMage)
	if err != nil {
		//TODO... 记录日志 抛出异常
	}
	yaml.Unmarshal(c0, &mageConfig)

	path := "./config/env_test.yaml"

	c, err := ioutil.ReadFile(path)
	if err != nil {
		//TODO... 记录日志 抛出异常
	}
	yaml.Unmarshal(c, &config)
	config.Mage = *mageConfig
	return
}
