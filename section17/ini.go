package main

import (
	"fmt"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("./section17/config.ini")
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustInt(8080),

		DbName: cfg.Section("db").Key("name").MustString("example.sql"),

		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main3() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
}
