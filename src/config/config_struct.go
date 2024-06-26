package config

type Configuration struct {
	// server listen address
	ServerAddr string `yaml:"ServerAddr"`

	MinVersion uint32 `yaml:"MinVersion"`

	// wechat config
	WXMini WechatConfig `yaml:"WXMini"`

	// qiniu config
	Qiniu QiniuConfig `yaml:"Qiniu"`

	// DBType: 1: SQLite3   2: MySQL
	DBType uint32 `yaml:"DBType"`

	// mysql config
	MySQL DBConfigMySQL `yaml:"MySQL"`

	// SQLite config
	SQLite DBConfigSQLite `yaml:"SQLite"`
}

// wechat config
type WechatConfig struct {
	AppID  string `yaml:"AppID"`
	Secret string `yaml:"Secret"`
}

// qiniu config
type QiniuConfig struct {
	Domain          string `yaml:"Domain"`
	Bucket          string `yaml:"Bucket"`
	AccessKey       string `yaml:"AccessKey"`
	SecretKey       string `yaml:"SecretKey"`
	UploadExpires   uint64 `yaml:"UploadExpires"`   // 上传凭证有效期, 单位: 分钟
	DownloadExpires uint64 `yaml:"DownloadExpires"` // 下载凭证有效期, 单位: 分钟
}

// mysql config
type DBConfigMySQL struct {
	DBHost     string `yaml:"DBHost"`
	DBUser     string `yaml:"DBUser"`
	DBUserPwd  string `yaml:"DBUserPwd"`
	DBDatabase string `yaml:"DBDatabase"`
	MaxConn    int    `yaml:"MaxConn"`
}

// SQLite config
type DBConfigSQLite struct {
	DBPath string `yaml:"DBPath"`
}

// ====================================================================

func (cfg *Configuration) IsValid() error {
	if err := cfg.MySQL.IsValid(); err != nil {
		return err
	}

	return nil
}

func (dbcfg *DBConfigMySQL) IsValid() error {
	return nil
}
