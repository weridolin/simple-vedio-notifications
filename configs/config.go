package configs

import (
	"os"
	"time"
)

var ConfigInstance *Config

type Config struct {
	PlatFormWhiteList              []string
	DBUri                          string
	NotifierBrokerUri              string
	DefaultTickerMaxSchedulerCount int
	DefaultMaxTickerCount          int
}

func GetAppConfig() *Config {
	time.Sleep(time.Second * 5)
	if ConfigInstance == nil {
		ConfigInstance = &Config{
			PlatFormWhiteList:              make([]string, 0),
			DBUri:                          os.Getenv("DBUri"),
			NotifierBrokerUri:              os.Getenv("NotifierBrokerUri"),
			DefaultTickerMaxSchedulerCount: 100,
			DefaultMaxTickerCount:          2,
		}
	}
	return ConfigInstance
}
