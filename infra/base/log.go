package base

import (
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

func init() {
	// 定义日志格式
	formatter := &prefixed.TextFormatter{}
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02 15:02:05"
	formatter.ForceFormatting = true // 强制格式化输出
	log.SetFormatter(formatter)
	// 日志级别
	level := os.Getenv("log.level")
	if level == "info" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}
	// 控制台高亮
	formatter.ForceColors = true
	// 日志文件
}
