package config

import (
	"fmt"
	"os"
)

type Mongo struct {
	Addr     string `json:"addr"`
	DB       string `json:"db"`
	AuthDB   string `json:"authDb"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Conf struct {
	Debug   bool   `json:"debug"`
	Trace   bool   `json:"trace"`
	LogPath string `json:"logPath"`
	Mongo   Mongo
}

var Config = &Conf{}

func init() {
	Config.Debug = getEnv("Debug", "true") == "true"
	Config.Trace = getEnv("Trace", "false") == "true"
	Config.LogPath = getEnv("LogPath", "/var/cainiaofund/log/%Y%m%d%H.log")

	Config.Mongo = Mongo{
		Addr:     getEnv("FMongoHost", ""),
		DB:       getEnv("FDB", ""),
		AuthDB:   getEnv("FAuthDb", ""),
		User:     getEnv("FUser", ""),
		Password: getEnv("FPassword", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultVal
}

func GetServerPort() string {
	return fmt.Sprintf(":%s", getEnv("FPort", "8081"))
}
