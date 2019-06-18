package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	PageSize        int
	RuntimeRootPath string
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
	TimeFormat      string
}

var AppSetting = &App{
	PageSize:    10,
	LogSavePath: "logs/",
	LogSaveName: "openstack-admin",
	LogFileExt:  "log",
	TimeFormat:  "20060102",
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{
	RunMode:      "debug",
	ReadTimeout:  60,
	WriteTimeout: 60,
	HttpPort:     8000,
}

type Nova struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var NovaSetting = &Nova{
	Type:     "mysql",
	Host:     "127.0.0.1",
	Port:     "3306",
	User:     "root",
	Password: "root",
	Name:     "nova",
}

type Redis struct {
	Host     string
	Port     string
	Password string
	MaxIdle  int
}

var RedisSetting = &Redis{
	Host:    "127.0.0.1",
	Port:    "6379",
	MaxIdle: 30,
}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("nova", NovaSetting)
	mapTo("redis", RedisSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
