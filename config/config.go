package config

import "os"

var (
	Debug   bool
	Trace   bool
	LogPath string
)

func init() {
	Debug = getEnv("Debug", true).(bool)
	Trace = getEnv("Trace", false).(bool)
	LogPath = "/var/cainiaofund/log/%Y%m%d%H%M.log"
}

func getEnv(key string, defaultVal interface{}) interface{} {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultVal
}
