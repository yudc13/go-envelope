package envelope

import (
	_ "goEnvelope/apis/web"
	_ "goEnvelope/core/account"
	"goEnvelope/infra"
	"goEnvelope/infra/base"
)

func init() {
	// 注册启动器
	// 配置文件
	infra.Register(&base.PropsStarter{})
	// 初始化数据库连接
	infra.Register(&base.DBStarter{})
	// 注册验证器
	infra.Register(&base.ValidateStarter{})
	// 注册iris
	infra.Register(&base.IrisStarter{})
	infra.Register(&infra.WebApiStarter{})
}
