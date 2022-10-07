package main

import (
	"fmt"

	"sample/api/db"
	"sample/api/models"
	"sample/api/router"
	"sample/api/setting"
)

func main() {
	//viper
	p, err := setting.ReadConfig()
	if err != nil {
		fmt.Println("read config error")
		return
	}

	//mongo
	m := db.Mongodb{}
	//帶入參數是讀取yaml進來的
	m.Connection(p.MongoHost, p.MongoUserName, p.MongoPassword)

	//寫入mongo
	m.Insert(models.Account{
		AccountID: 123456,
		Limit:     7890,
		Products:  nil,
	})

	//讀取
	data := m.ReadALL("test", "accounts")
	fmt.Println("=====================================")
	fmt.Println(data)
	fmt.Println("=====================================")

	//gin
	router.StartService()
}
