package main

import (
	"fmt"
	"github.com/tietang/props/v3/ini"
)

func main() {
	source := ini.NewIniFileConfigSource("example/config.ini")
	port := source.GetIntDefault("app.server.port", 3000)
	fmt.Println(port)
}
