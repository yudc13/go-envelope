package infra

import (
	log "github.com/sirupsen/logrus"
	"github.com/tietang/props/v3/kvs"
	"reflect"
	"sort"
)

const (
	KeyProps = "_conf"
)

// 资源启动器上下文
type StarterContext map[string]interface{}

type PriorityGroup int

type Starter interface {
	// 资源初始化
	Init(ctx StarterContext)
	// 资源安装
	Setup(ctx StarterContext)
	// 启动资源
	Start(ctx StarterContext)
	// 说明该资源启动器开始启动时是否会阻塞
	// 如果存在多个阻塞启动器时，只有阻塞最后一个，之前的会通过goroutine来异步启动
	// 所以需要规划好启动器注册顺序
	StartBlocking() bool
	// 资源停止
	Stop(ctx StarterContext)
	PriorityGroup() PriorityGroup
	Priority() int
}

// 服务启动注册器
type starterRegister struct {
	nonBlockingStarters []Starter
	blockingStarters    []Starter
}

func (s StarterContext) Props() kvs.ConfigSource {
	p := s[KeyProps]
	if p == nil {
		panic("配置还没有被初始化")
	}
	return p.(kvs.ConfigSource)
}

func (s StarterContext) SetProps(conf kvs.ConfigSource) {
	s[KeyProps] = conf
}

// 返回所有启动器
func (r *starterRegister) AllStarters() []Starter {
	starters := make([]Starter, 0)
	starters = append(starters, r.blockingStarters...)
	starters = append(starters, r.nonBlockingStarters...)
	return starters
}

// 注册启动器
func (r *starterRegister) Register(s Starter) {
	if s.StartBlocking() {
		r.blockingStarters = append(r.blockingStarters, s)
	} else {
		r.nonBlockingStarters = append(r.nonBlockingStarters, s)
	}
	t := reflect.TypeOf(s)
	log.Infof("Register starter: %s \n", t.String())
}

var StarterRegister *starterRegister = &starterRegister{}

type Starters []Starter

func (s Starters) Len() int {
	return len(s)
}
func (s Starters) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Starters) Less(i, j int) bool {
	return s[i].PriorityGroup() > s[j].PriorityGroup() && s[i].Priority() > s[j].Priority()
}

// 注册starter
func Register(s Starter) {
	StarterRegister.Register(s)
}
func SortStarters() {
	sort.Sort(Starters(StarterRegister.AllStarters()))
}

// 获取所有注册的starter
func GetStarters() []Starter {
	return StarterRegister.AllStarters()
}

const (
	SystemGroup         PriorityGroup = 30
	BasicResourcesGroup PriorityGroup = 20
	AppGroup            PriorityGroup = 10

	DefaultPriority = 10000
)

// 默认空实现 方便资源启动器的实现
type BaseStarter struct{}

func (s *BaseStarter) Init(ctx StarterContext)      {}
func (s *BaseStarter) Setup(ctx StarterContext)     {}
func (s *BaseStarter) Start(ctx StarterContext)     {}
func (s *BaseStarter) Stop(ctx StarterContext)      {}
func (s *BaseStarter) StartBlocking() bool          { return false }
func (s *BaseStarter) PriorityGroup() PriorityGroup { return BasicResourcesGroup }
func (s *BaseStarter) Priority() int                { return DefaultPriority }
