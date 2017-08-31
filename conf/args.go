package conf

import (
	"github.com/spf13/viper"
	"github.com/sirupsen/logrus"
)

var Args Arguments

type Arguments struct {
	CouchbaseServers string `mapstructure:"couchbase_servers"`
	LogLevel         string `mapstructure:"log_level"`
}

func init() {
	setDefaults()
	v := viper.New()
	v.SetConfigName("rima") // name of config file (without extension)
	v.AddConfigPath(".")    // optionally look for config in the working directory
	v.AddConfigPath("$GOPATH/bin")
	v.AddConfigPath("$HOME")
	v.AddConfigPath("$HOME/go/bin")
	err := v.ReadInConfig()
	if err != nil {
		logrus.Errorf("config file error: %+v", err)
		return
	}
	err = v.Unmarshal(&Args) // Find and read the config file
	if err != nil {
		logrus.Infof("config file error: %s", err)
	} else {
		logrus.Infof("Configuration: %+v", Args)
		//v.WatchConfig()
	}
}

func setDefaults() {
	Args.LogLevel = "info"
}