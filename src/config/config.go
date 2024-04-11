package config

import (
	"flag"
	"strings"

	"github.com/golang/glog"
	"github.com/lizisky/liziutils/utils"
)

var (
	gConfig  *Configuration
	gDatadir = flag.String("data_dir", "", "data directory")
)

func GetConfig() *Configuration {
	return gConfig
}

func GetDataDir() string {
	return *gDatadir
}

func LoadConfig() bool {
	const (
		data_root_dir   = ".lizi_data"
		config_fileName = "lizisky_cfg.yml"
	)

	config_tmp := Configuration{}
	general_data_dir, err := utils.LoadConfigurationFromYamlFile(*gDatadir, data_root_dir, config_fileName, &config_tmp)
	if err != nil {
		panic(err.Error())
	}

	if strings.Compare(*gDatadir, general_data_dir) != 0 {
		gDatadir = &general_data_dir
	}

	gConfig = &config_tmp

	glog.Info("Success Load configuration from directory:", *gDatadir, "\n", utils.ToJSONIndent(config_tmp))

	return true
}
