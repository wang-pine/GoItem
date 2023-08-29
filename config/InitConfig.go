package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var ConfigInfo *viper.Viper

// config结构体，用于返回对应得config信息
type Config struct {
	LocalAddr string `json:"local_addr"`
	MysqlIP   string `json:"mysql_ip"`
	MysqlPort string `json:"mysql_port"`
}

// 热配置函数
func InitConfig() {
	ConfigInfo = initConfig()
	go dynamicConfig()
}
func initConfig() *viper.Viper {
	//v := viper.New()
	v := viper.New()
	v.SetConfigName("config")   // 设置文件名称（无后缀）
	v.SetConfigType("yaml")     // 设置后缀名 {"1.6以后的版本可以不设置该后缀"}
	v.AddConfigPath("./config") // 设置文件所在路径
	v.Set("verbose", true)      // 设置默认参数

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(" Config file not found; ignore error if desired")
		} else {
			panic("Config file was found but another error was produced")
		}
	}
	return v
}
func dynamicConfig() {
	// 监控配置和重新获取配置
	ConfigInfo.WatchConfig()
	ConfigInfo.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}


func GetConfig() (config Config) {
	fmt.Println("正在获取配置文件...")
	url := "http://localhost:8080/douyin/config"
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	jsonErr := json.Unmarshal(body, &config)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	return config
	// fmt.Println("配置罗列")
	// fmt.Println("local_addr =", config.LocalAddr)
	// fmt.Println("mysql_ip=", config.MysqlIP)
	// fmt.Println("mysql_port=", config.MysqlPort)
}
func GetDBAddr() (dbAddr string) {
	config := GetConfig()
	dbAddr = config.MysqlIP + ":" + config.MysqlPort
	return dbAddr
}
func GetLocalAddr() string {
	config := GetConfig()
	return config.LocalAddr
}
