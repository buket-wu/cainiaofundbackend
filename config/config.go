package config

import "os"

type Conf struct {
	Debug   bool   `json:"debug"`
	Trace   bool   `json:"trace"`
	LogPath string `json:"logPath"`
}

var Config = &Conf{}

func init() {
	Config.Debug = getEnv("Debug", true).(bool)
	Config.Trace = getEnv("Trace", false).(bool)
	Config.LogPath = getEnv("LogPath", "/var/cainiaofund/log/%Y%m%d%H.log").(string)
}

func getEnv(key string, defaultVal interface{}) interface{} {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultVal
}

func GetServerPort() string {
	return ":8080"
}
