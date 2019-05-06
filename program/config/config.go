package config

import (
	"os"

	"github.com/etcd-manage/etcd-manage-server/program/common"
	"github.com/naoina/toml"
)

// Config 配置
type Config struct {
	Debug   bool         `toml:"debug"`
	LogPath string       `toml:"log_path"`
	HTTP    *HTTP        `toml:"http"`
	DB      *MySQLConfig `toml:"db"`
}

// HTTP http 件套配置
type HTTP struct {
	Address               string   `toml:"address"`
	Port                  int      `toml:"port"`
	TLSEnable             bool     `toml:"tls_enable"`               // 是否启用tls连接
	TLSConfig             *HTTPTls `toml:"tls_config"`               // 启用tls时必须配置此内容
	TLSEncryptEnable      bool     `toml:"tls_encrypt_enable"`       // 是否启用 Let's Encrypt tls
	TLSEncryptDomainNames []string `toml:"tls_encrypt_domain_names"` // 启用 Let's Encrypt 时的域名列表
}

// HTTPTls http tls配置
type HTTPTls struct {
	CertFile string `toml:"cert_file"`
	KeyFile  string `toml:"key_file"`
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	Debug        bool   `toml:"debug"`          // 是否调试模式
	Address      string `toml:"address"`        // 数据库连接地址
	Port         int    `toml:"port"`           // 数据库端口
	MaxIdleConns int    `toml:"max_idle_conns"` // 连接池最大连接数
	MaxOpenConns int    `toml:"max_open_conns"` // 默认打开连接数
	User         string `toml:"user"`           // 数据库用户名
	Passwd       string `toml:"passwd"`         // 数据库密码
	DbName       string `toml:"db_name"`        // 数据库名
}

var (
	cfg *Config
)

// GetCfg 获取配置
func GetCfg() *Config {
	if cfg == nil {
		LoadConfig("")
	}
	return cfg
}

// LoadConfig 读取配置
func LoadConfig(cfgPath string) (*Config, error) {
	cfgPath = getCfgPath(cfgPath)
	f, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	cfg = new(Config)
	if err := toml.NewDecoder(f).Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func getCfgPath(cfgPath string) string {
	if cfgPath == "" {
		cfgPath = common.GetRootDir() + "config/cfg.toml"
	}
	return cfgPath
}
