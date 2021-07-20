package config

import "os"

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
	Config.Debug = getEnv("Debug", true) == "1"
	Config.Trace = getEnv("Trace", false) == "1"
	Config.LogPath = getEnv("LogPath", "/var/cainiaofund/log/%Y%m%d%H.log").(string)

	mongoJson := getEnv("mongo", "")
	if mongoJson == "" {
		Config.Mongo = Mongo{
			Addr:     "127.0.0.1:27017",
			DB:       "cainiaofund",
			AuthDB:   "cainiaofund",
			User:     "cainiao",
			Password: "qweasdzxc",
		}
	}
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
