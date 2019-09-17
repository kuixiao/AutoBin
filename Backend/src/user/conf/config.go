package conf

import (
	"encoding/json"
	"fmt"
	io "io/ioutil"
	"log"
	"os"
	"sync"
)

//定义配置文件解析后的结构
type ConfigInfo struct {
	UserName  		string `json:userName`
	Password  		string `json:password`
	Host      		string `json:host`
	Port  	  		string `json:port`
	DbName    		string `json:dbName`
	GrpcPort		string `json:grpcPort`
	HttpPort		string `json:httpPort`
	Verify			string `json:verify`
	Prefix			string `json:prefix`
	GrpcEndpoint	string `json:grpcEndpoint`
}

var Config ConfigInfo
var file_locker sync.Mutex //文件锁

// 初始化文件配置
func init()  {
	conf, err := LoadConfig("\\config.json") //get config struct
	if err != nil {
		log.Fatal("InitConfig failed:", err.Error())
	}
	Config = conf
}

// 加载配置文件
func LoadConfig(filename string) (ConfigInfo, error) {
	var conf ConfigInfo
	file_locker.Lock()
	data, err := io.ReadFile(getCurrentPath()+filename) //read config file
	file_locker.Unlock()
	if err != nil {
		fmt.Println("read json file error")
		return conf, err
	}
	datajson := []byte(data)
	err = json.Unmarshal(datajson, &conf)
	if err != nil {
		fmt.Println("unmarshal json file error")
		return conf, err
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