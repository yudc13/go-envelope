package main

import (
	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func main() {
	app := iris.New()
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(map[string]string{
			"name": "jack",
		})
	})
	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		id := ctx.Params().GetUint64Default("id", 1)
		ctx.WriteString(strconv.Itoa(int(id)))
	})
	err := app.Listen(":8888")
	log.Fatal(err)
}
