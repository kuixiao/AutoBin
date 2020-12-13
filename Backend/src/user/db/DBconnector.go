package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kuixiao/AutoBin/Backend/src/user/conf"
	"log"
)

type DBconfig struct{
	User 		string
	Password	string
	Host		string
	Port		string
	DBname		string
}

var DB *gorm.DB

func init(){
	//var config = DBconfig{
	//	"root",
	//	"123456",
	//	"127.0.0.1",
	//	"3306",
	//	"user",
	//}
	config := conf.Config
	var err error
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.UserName,config.Password, config.Host, config.Port,config.DbName )
	log.Printf("connArgs: %s", connArgs)
	DB, err = gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatalf("open mysql err: %s", err.Error())
	}
	log.Println()
}
