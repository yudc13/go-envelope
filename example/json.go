package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

type User struct {
	Id      int    `json:"id,string"` // string 序列化后的值转为string类型
	Name    string `json:"username"`
	Age     int    `json:"age,omitempty"` // omitempty 忽律零值 - 绝对忽律
	Address string `json:"address"`
}

func main() {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	user := &User{
		Id:      1,
		Name:    "jack",
		Age:     0,
		Address: "北京",
	}
	// 序列化
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b)) // {"id":1,"name":"jack","age":23,"address":"北京"}
}
