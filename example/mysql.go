package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	dsName := "root:yudachao@tcp(127.0.0.1:3306)/book?charset=utf8&parseTime=true&loc=Local"
	db, err := sql.Open("mysql", dsName)
	if err != nil {
		log.Fatal(err)
	}
	// 设置空闲连接池中连接的最大数量
	db.SetMaxIdleConns(2)
	// 设置打开数据库连接的最大数量
	db.SetMaxOpenConns(3)
	// 设置连接可复用的最大时间
	db.SetConnMaxLifetime(7 * time.Hour)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mysql connected success")
}
