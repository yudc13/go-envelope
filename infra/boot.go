package infra

import (
	"github.com/tietang/props/v3/kvs"
)

type BootApplication struct {
	// 配置
	conf           kvs.ConfigSource
	starterContext StarterContext
}

func New(conf kvs.ConfigSource) *BootApplication {
	b := &BootApplication{
		conf:           conf,
		starterContext: StarterContext{},
	}
	b.starterContext.SetProps(conf)
	return b
}

func (b *BootApplication) Start() {
	// 1. 初始化所有的Starter
	b.init()
	// 2. 安装所有的Starter
	b.setup()
	// 3. 启动Starter
	b.start()
}

func (b *BootApplication) init() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(b.starterContext)
	}
}

func (b *BootApplication) setup() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(b.starterContext)
	}
}

func (b *BootApplication) start() {
	allSTarters := StarterRegister.AllStarters()
	for i, starter := range allSTarters {
		if starter.StartBlocking() {
			// 如果可阻塞的是最后一个 直接启动
			if i+1 == len(allSTarters) {
				starter.Start(b.starterContext)
			} else {
				// 否则使用goroutine来异步启动
				go starter.Start(b.starterContext)
			}
		} else {
			starter.Start(b.starterContext)
		}
	}
}
