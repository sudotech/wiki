package main

import (
	"fmt"
	"net/http"

	"github.com/donnie4w/go-logger/logger"
	"gome.com.cn/httpserver/bootstrap"
)

func main() {

	fmt.Println("服务器启动！")
	bootstrap.InitLog()
	bootstrap.InitRouter()

	logger.Info("启动监听端口：", "8000")
	//配置说明
	http.ListenAndServe(":8000", nil)
	fmt.Println("站点停止！")

}
