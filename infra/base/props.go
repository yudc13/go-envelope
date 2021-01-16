package base

import (
	log "github.com/sirupsen/logrus"
	"github.com/tietang/props/v3/kvs"
	"goEnvelope/infra"
)

var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	return props
}

type PropsStarter struct {
	infra.BaseStarter
}

func (p *PropsStarter) Init(ctx infra.StarterContext) {
	props = ctx.Props()
	log.Info("Props init success")
}
