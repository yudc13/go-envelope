package base

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	irisRecover "github.com/kataras/iris/v12/middleware/recover"
	"github.com/sirupsen/logrus"
	"goEnvelope/infra"
	"time"
)

var irisApplication *iris.Application

func Iris() *iris.Application {
	return irisApplication
}

type IrisStarter struct {
	infra.BaseStarter
}

func (i *IrisStarter) Init(ctx infra.StarterContext) {
	irisApplication = initIris()
	log := irisApplication.Logger()
	log.Install(logrus.StandardLogger())
}

func (i *IrisStarter) Start(ctx infra.StarterContext) {
	routes := Iris().GetRoutes()
	for _, r := range routes {
		logrus.Infof("method: %s path: %s", r.Method, r.Path)
	}
	// 启动
	port := ctx.Props().GetDefault("app.server.port", "8080")
	err := Iris().Run(iris.Addr(":" + port))
	if err != nil {
		logrus.Errorf("Iris start error: %v \n", err)
	}
	logrus.Infof("Iris started \n")
}

// 是阻塞的
func (i *IrisStarter) StartBlocking() bool {
	return false
}

// 初始化iris
func initIris() *iris.Application {
	app := iris.New()
	app.Use(irisRecover.New())
	cfg := logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
		Query:  true,
		LogFunc: func(
			endTime time.Time,
			latency time.Duration,
			status, ip, method, path string,
			message interface{}, headerMessage interface{}) {
			app.Logger().Infof("| %s | %s | %s | %s | %s | %s | %s |",
				endTime.Format("2006-01-02 15:04:05"),
				latency.String(),
				status,
				ip,
				method,
				message,
				headerMessage,
			)
		},
	}
	app.Use(logger.New(cfg))
	return app
}

