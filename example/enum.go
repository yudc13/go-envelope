package main

import "fmt"

// golang中的"枚举常量"

const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Staturday = 6
)

// 隐式复制前一个非空表达式
const (
	Apply, Banana     = 11, 22
	Strawberry, Grape // 11 22
	Pear, Watermelon  // 11 22
)

// iota 位于同一行的iota 即便出现多次 其值也是一样的
const (
	a, b = iota, iota + 10 // 0, 10 (iota = 0)
	c, d                   // 1, 11 (iota = 1)
	e, f                   // 2 12 (iota = 2)
)

// 略过 iota = 0
const (
	_ = iota
	g // 1
	h // 2
)

// 使用有类型的枚举常量保证类型安全
// 枚举常量多数也是无类型常量
type myInt int32

const (
	i myInt = iota
)

func main() {
	j := i + h
	fmt.Println(g, h, j)
	fmt.Printf("h: %T \n", h) // int
}
