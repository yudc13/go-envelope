package infra

type Initializer interface {
	// 用户对象实例化后的初始化操作
	Init()
}

// 初始化注册器
type InitailizeRegister struct {
	Initializers []Initializer
}

func (i *InitailizeRegister) Register(a Initializer)  {
	i.Initializers = append(i.Initializers, a)
}
