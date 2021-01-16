package base

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"goEnvelope/infra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库实例
var db *gorm.DB

func DB() *gorm.DB {
	return db
}

// 数据库starter
type DBStarter struct {
	infra.BaseStarter
}

func (s *DBStarter) Init(ctx infra.StarterContext) {
	conf := ctx.Props()
	username := conf.GetDefault("mysql.username", "root")
	password := conf.GetDefault("mysql.password", "yudachao")
	host := conf.GetDefault("mysql.host", "127.0.0.1:3306")
	databaseName := conf.GetDefault("mysql.database", "envelope")
	database, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", username, password, host, databaseName),
		SkipInitializeWithVersion: false, // 根据当前mysql版本自动配置
		DefaultStringSize:         256,   // string类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用datetime精度 mysql5.6之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并创建的方式 MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	}), &gorm.Config{})
	if err != nil {
		log.Errorf("DB init failed: %v", err)
		panic(err)
	}
	db = database
	log.Info("DB init success")
}
