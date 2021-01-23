package infra

var apiInitailizerRegister *InitailizeRegister = &InitailizeRegister{}

// 注册web api初始化对象
func RegisterApi(a Initializer)  {
	apiInitailizerRegister.Register(a)
}

func GetApiInitailizers() []Initializer {
	return apiInitailizerRegister.Initializers
}

type WebApiStarter struct {
	BaseStarter
}

func (w *WebApiStarter) Setup(ctx StarterContext)  {
	for _, v := range GetApiInitailizers() {
		v.Init()
	}
}

