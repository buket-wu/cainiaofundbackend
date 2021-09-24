package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Conf struct {
	Debug      bool   `yaml:"Debug"`      //debug 暂时只用于日记
	Trace      bool   `yaml:"Trace"`      // trace 暂时只用于日记
	LogPath    string `yaml:"LogPath"`    // 日记保存路径
	Mongo      Mongo  `yaml:"Mongo"`      // mongo配置
	RoutineNum int    `yaml:"RoutineNum"` // 协程池容量
}

type Mongo struct {
	Addr     string `yaml:"Addr"`
	DB       string `yaml:"DB"`
	AuthDB   string `yaml:"AuthDB"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
}

var Config = &Conf{}

func init() {
	confYamlPath := getEnv("CainiaofundConfigPath", "")
	if confYamlPath == "" {
		log.Fatal("invalid yamlConfPath")
	}

	f, err := os.Open(confYamlPath)
	if err != nil {
		log.Fatalf("open yaml file fail; err:%v", err)
	}
	yamlContent, _ := ioutil.ReadAll(f)

	err = yaml.Unmarshal(yamlContent, Config)
	if err != nil {
		log.Fatalf("unmarshal conf yaml fail; err:%v", err)
	}
}

func getEnv(key string, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultVal
}

func GetServerPort() string {
	return fmt.Sprintf(":%s", getEnv("FPort", "8090"))
}
