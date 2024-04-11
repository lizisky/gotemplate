package config

type Configuration struct {
	// server listen address
	ServerAddr string `yaml:"ServerAddr"`

	MinVersion uint32 `yaml:"MinVersion"`

	// wechat config
	WXMini WechatConfig `yaml:"WXMini"`

	// qiniu config
	Qiniu QiniuConfig `yaml:"Qiniu"`

	// mysql config
	MySQL DBConfig `yaml:"MySQL"`
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
type DBConfig struct {
	DBHost     string `yaml:"DBHost"`
	DBUser     string `yaml:"DBUser"`
	DBUserPwd  string `yaml:"DBUserPwd"`
	DBDatabase string `yaml:"DBDatabase"`
	MaxConn    int    `yaml:"MaxConn"`
}

func (cfg *Configuration) IsValid() error {
	if err := cfg.MySQL.IsValid(); err != nil {
		return err
	}

	return nil
}

func (dbcfg *DBConfig) IsValid() error {
	return nil
}
