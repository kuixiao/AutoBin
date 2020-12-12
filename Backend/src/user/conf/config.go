package conf

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"os"
)

//定义配置文件解析后的结构
type ConfigInfo struct {
	ConfigPath		string

	UserName  		string `json:"userName"`
	Password  		string `json:"password"`
	Host      		string `json:"host"`
	Port  	  		string `json:"port"`
	DbName    		string `json:"dbName"`
	GrpcPort		string `json:"grpcPort"`
	HttpPort		string `json:"httpPort"`
	Verify			string `json:"verify"`
	Prefix			string `json:"prefix"`
	GrpcEndpoint	string `json:"grpcEndpoint"`
	HttpHost		string `json:"httpHost"`

	AppKey			string `json:"appKey"`		// 百度云应用键值
	Secret			string `json:"secret"`		// 百度云应用秘钥
	AccessToken		string `json:"accessToken"`	// 调api使用的token
	ApiUrl			string `json:"apiUrl"`		// api路径 获取物品名
	ClassApiUrl		string `json:"ClassApiUrl"`	// api路径 获取分类
}

var Config ConfigInfo

// 初始化文件配置
func init()  {
	conf, err := LoadConfig() //get config struct
	if err != nil {
		log.Fatal("InitConfig failed:", err.Error())
	}
	Config = conf
}

func (config *ConfigInfo) parseFlag() {
	flag.StringVar(&config.ConfigPath, "config", "", "config file path")
	flag.Parse()
}

// 加载配置文件
func LoadConfig() (ConfigInfo, error) {
	var conf ConfigInfo
	conf.parseFlag()
	log.Printf("conf.ConfigPath: %s", conf.ConfigPath)
	viper.SetConfigFile(conf.ConfigPath)
	if err := viper.ReadInConfig(); err != nil {
		return ConfigInfo{}, err
	}
	if err := viper.Unmarshal(&conf); err != nil {
		return ConfigInfo{}, err
	}
	return conf, nil
}


//获取当前路径，比如：E:/abc/data/test
func getCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir += "\\Backend\\src\\user\\conf"
	return dir
}