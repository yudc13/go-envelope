package main

import (
	"github.com/tietang/props/v3/ini"
	_ "goEnvelope"
	"goEnvelope/infra"
)

func main() {
	// 加载解析配置文件
	conf := ini.NewIniFileCompositeConfigSource("example/config.ini")
	app := infra.New(conf)
	app.Start()
}
