package testx

import (
	"github.com/tietang/props/v3/ini"
	"goEnvelope/infra"
	"goEnvelope/infra/base"
)

func init() {
	// 加载解析配置文件
	conf := ini.NewIniFileCompositeConfigSource("../../brun/config.ini")
	// 配置文件
	infra.Register(&base.PropsStarter{})
	// 初始化数据库连接
	infra.Register(&base.DBStarter{})
	// 注册验证器
	infra.Register(&base.ValidateStarter{})
	app := infra.New(conf)
	app.Start()
}
