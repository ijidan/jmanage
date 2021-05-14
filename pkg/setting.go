package pkg

import (
	"gopkg.in/ini.v1"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	PageSize  int
	JwtSecret string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)

//初始化
func Init() {
	LoadCfg()
	loadBase()
	loadApp()
	loadServer()
}

//加载配置文件
func LoadCfg() {
	var err error
	Cfg, err = ini.Load("my.ini")
	if err != nil {
		panic(err)
	}
}

//加载基础配置
func loadBase() {
	section := Cfg.Section("")
	RunMode = section.Key("app_mode").String()
}

//加载APP配置
func loadApp() {
	section := Cfg.Section("app")
	PageSize, _ = section.Key("page_size").Int()
	JwtSecret = section.Key("jwt_secret").String()
}

//加载server配置
func loadServer() {
	section := Cfg.Section("server")
	HTTPPort, _ = section.Key("http_port").Int()
	ReadTimeoutInt, _ := section.Key("read_timeout").Int64()
	WriteTimeoutInt, _ := section.Key("write_timeout").Int64()

	ReadTimeout = time.Duration(ReadTimeoutInt)
	WriteTimeout = time.Duration(WriteTimeoutInt)

}
