package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/daymenu/gostudy/examples/ent/biz"

	"github.com/daymenu/gostudy/examples/ent/ent"
	// 初始化数据库
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ent?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	timeCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	u, err := biz.CreateUser(timeCtx, client)
	if err != nil {
		fmt.Println("error")
	}

	fmt.Println(u)
}
